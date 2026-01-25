// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IPoll} from "./interfaces/IPoll.sol";
import {IZKVerifier} from "./interfaces/IZKVerifier.sol";
import {IOracle} from "./interfaces/IOracle.sol";

/**
 * @title Poll
 * @notice Confidential poll with commit-reveal voting and ZK proofs
 * @dev Implements privacy-preserving voting with:
 *      - Commit phase: voters submit commitments with ZK proofs
 *      - Reveal phase: voters reveal their votes after poll closes
 *      - Tally phase: final results are computed
 */
contract Poll is IPoll {
    /// @notice Current state of the poll
    PollState public override state;

    /// @notice Poll question
    string public override question;

    /// @notice Available options
    string[] public override options;

    /// @notice Poll creator
    address public immutable creator;

    /// @notice Poll creation timestamp
    uint256 public immutable createdAt;

    /// @notice Poll duration in seconds
    uint256 public immutable duration;

    /// @notice Poll end time
    uint256 public immutable endTime;

    /// @notice Merkle root of eligible voters
    bytes32 public immutable voterMerkleRoot;

    /// @notice ZK proof verifier
    IZKVerifier public immutable zkVerifier;

    /// @notice Oracle for poll closing
    IOracle public immutable oracle;

    /// @notice Mapping of voter address to their commitment
    mapping(address => VoteCommitment) public commitments;

    /// @notice Tallied results (index = choice, value = count)
    uint256[] private _results;

    /// @notice Total number of votes committed
    uint256 public totalCommitted;

    /// @notice Total number of votes revealed
    uint256 public totalRevealed;

    /// @notice Custom errors
    error PollNotActive();
    error PollNotClosed();
    error PollNotTallied();
    error PollAlreadyClosed();
    error PollAlreadyTallied();
    error UnauthorizedOracle();
    error AlreadyVoted();
    error InvalidChoice();
    error InvalidProof();
    error InvalidMerkleProof();
    error NoCommitment();
    error AlreadyRevealed();
    error InvalidReveal();

    /**
     * @notice Create a new poll
     * @param _question The poll question
     * @param _options Array of voting options
     * @param _duration Poll duration in seconds
     * @param _voterMerkleRoot Merkle root of eligible voters
     * @param _zkVerifier Address of ZK proof verifier
     * @param _oracle Address of oracle for poll closing
     */
    constructor(
        string memory _question,
        string[] memory _options,
        uint256 _duration,
        bytes32 _voterMerkleRoot,
        address _zkVerifier,
        address _oracle
    ) {
        require(_options.length >= 2, "At least 2 options required");
        require(_duration > 0, "Duration must be positive");
        require(_voterMerkleRoot != bytes32(0), "Invalid Merkle root");
        require(_zkVerifier != address(0), "Invalid verifier");
        require(_oracle != address(0), "Invalid oracle");

        creator = msg.sender;
        question = _question;
        options = _options;
        duration = _duration;
        voterMerkleRoot = _voterMerkleRoot;
        zkVerifier = IZKVerifier(_zkVerifier);
        oracle = IOracle(_oracle);

        createdAt = block.timestamp;
        endTime = block.timestamp + _duration;
        state = PollState.Active;

        // Initialize results array
        _results = new uint256[](_options.length);

        // Request oracle to close poll at end time
        oracle.requestPollClose(address(this), endTime);
    }

    /**
     * @notice Commit a vote with ZK proof of eligibility
     * @param commitment Hash of (choice, salt, voter)
     * @param zkProof ZK proof of voter eligibility and vote validity
     * @param merklePath Merkle proof of voter in eligible set
     */
    function commitVote(
        bytes32 commitment,
        bytes calldata zkProof,
        bytes32[] calldata merklePath
    ) external override {
        if (state != PollState.Active) revert PollNotActive();
        if (commitments[msg.sender].commitment != bytes32(0)) revert AlreadyVoted();

        // Verify Merkle proof of voter eligibility
        if (!_verifyMerkleProof(msg.sender, merklePath)) {
            revert InvalidMerkleProof();
        }

        // Verify ZK proof (proof includes commitment as public input)
        uint256[] memory publicInputs = new uint256[](1);
        publicInputs[0] = uint256(commitment);

        if (!zkVerifier.verifyProof(zkProof, publicInputs)) {
            revert InvalidProof();
        }

        // Store commitment
        commitments[msg.sender] = VoteCommitment({
            commitment: commitment,
            timestamp: block.timestamp,
            revealed: false
        });

        totalCommitted++;

        emit VoteCommitted(msg.sender, commitment, block.timestamp);
    }

    /**
     * @notice Reveal a committed vote
     * @param choice The vote choice (index in options array)
     * @param salt Random salt used in commitment
     */
    function revealVote(uint256 choice, bytes32 salt) external override {
        if (state != PollState.Closed) revert PollNotClosed();

        VoteCommitment storage commitment = commitments[msg.sender];

        if (commitment.commitment == bytes32(0)) revert NoCommitment();
        if (commitment.revealed) revert AlreadyRevealed();
        if (choice >= options.length) revert InvalidChoice();

        // Verify reveal matches commitment
        bytes32 expectedCommitment = keccak256(abi.encodePacked(choice, salt, msg.sender));

        if (expectedCommitment != commitment.commitment) {
            revert InvalidReveal();
        }

        // Mark as revealed and update tally
        commitment.revealed = true;
        _results[choice]++;
        totalRevealed++;

        emit VoteRevealed(msg.sender, choice, block.timestamp);
    }

    /**
     * @notice Close the poll (only callable by oracle)
     */
    function closePoll() external override {
        if (msg.sender != address(oracle)) revert UnauthorizedOracle();
        if (state != PollState.Active) revert PollAlreadyClosed();

        state = PollState.Closed;

        emit PollClosed(block.timestamp);
    }

    /**
     * @notice Tally final results (can be called by anyone after closing)
     */
    function tally() external override returns (uint256[] memory) {
        if (state != PollState.Closed) revert PollNotClosed();

        state = PollState.Tallied;

        emit ResultsTallied(_results, block.timestamp);

        return _results;
    }

    /**
     * @notice Get voter's commitment
     * @param voter Address of the voter
     */
    function getCommitment(address voter) external view override returns (VoteCommitment memory) {
        return commitments[voter];
    }

    /**
     * @notice Get tallied results
     */
    function getResults() external view override returns (uint256[] memory) {
        if (state != PollState.Tallied) revert PollNotTallied();
        return _results;
    }

    /**
     * @notice Verify Merkle proof of voter eligibility
     * @param voter Address of the voter
     * @param proof Merkle proof path
     */
    function _verifyMerkleProof(address voter, bytes32[] calldata proof) private view returns (bool) {
        bytes32 leaf = keccak256(abi.encodePacked(voter));
        bytes32 computedHash = leaf;

        for (uint256 i = 0; i < proof.length; i++) {
            bytes32 proofElement = proof[i];

            if (computedHash <= proofElement) {
                computedHash = keccak256(abi.encodePacked(computedHash, proofElement));
            } else {
                computedHash = keccak256(abi.encodePacked(proofElement, computedHash));
            }
        }

        return computedHash == voterMerkleRoot;
    }
}
