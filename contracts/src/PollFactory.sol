// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Poll} from "./Poll.sol";
import {IOracle} from "./interfaces/IOracle.sol";

/**
 * @title PollFactory
 * @notice Factory contract for creating and managing polls
 * @dev Deploys new Poll contracts and maintains registry
 */
contract PollFactory {
    /// @notice Address of the ZK verifier contract
    address public immutable zkVerifier;

    /// @notice Address of the oracle contract
    address public immutable oracle;

    /// @notice Counter for poll IDs
    uint256 public pollCount;

    /// @notice Mapping from poll ID to poll address
    mapping(uint256 => address) public polls;

    /// @notice Mapping from poll address to poll ID
    mapping(address => uint256) public pollIds;

    /// @notice Emitted when a new poll is created
    event PollCreated(
        uint256 indexed pollId,
        address indexed pollAddress,
        address indexed creator,
        string question,
        uint256 duration
    );

    /// @notice Custom errors
    error InvalidVerifier();
    error InvalidOracle();

    /**
     * @notice Constructor
     * @param _zkVerifier Address of ZK proof verifier
     * @param _oracle Address of oracle for poll closing
     */
    constructor(address _zkVerifier, address _oracle) {
        if (_zkVerifier == address(0)) revert InvalidVerifier();
        if (_oracle == address(0)) revert InvalidOracle();

        zkVerifier = _zkVerifier;
        oracle = _oracle;
    }

    /**
     * @notice Create a new poll
     * @param question The poll question
     * @param options Array of voting options (minimum 2)
     * @param duration Poll duration in seconds
     * @param voterMerkleRoot Merkle root of eligible voters
     * @return pollAddress Address of the newly created poll
     */
    function createPoll(
        string memory question,
        string[] memory options,
        uint256 duration,
        bytes32 voterMerkleRoot
    ) external returns (address pollAddress) {
        // Create new poll
        Poll poll = new Poll(
            question,
            options,
            duration,
            voterMerkleRoot,
            zkVerifier,
            oracle
        );

        pollAddress = address(poll);
        uint256 pollId = pollCount++;

        // Store in registry
        polls[pollId] = pollAddress;
        pollIds[pollAddress] = pollId;

        emit PollCreated(pollId, pollAddress, msg.sender, question, duration);

        return pollAddress;
    }

    /**
     * @notice Get poll address by ID
     * @param pollId The poll ID
     * @return Poll contract address
     */
    function getPoll(uint256 pollId) external view returns (address) {
        return polls[pollId];
    }

    /**
     * @notice Get poll ID by address
     * @param pollAddress The poll contract address
     * @return Poll ID
     */
    function getPollId(address pollAddress) external view returns (uint256) {
        return pollIds[pollAddress];
    }

    /**
     * @notice Get total number of polls created
     * @return Total poll count
     */
    function getTotalPolls() external view returns (uint256) {
        return pollCount;
    }
}
