// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IPoll
 * @notice Interface for confidential poll contract
 */
interface IPoll {
    /// @notice Poll states
    enum PollState {
        Active,
        Closed,
        Tallied
    }

    /// @notice Vote commitment structure
    struct VoteCommitment {
        bytes32 commitment;
        uint256 timestamp;
        bool revealed;
    }

    /// @notice Emitted when a vote is committed
    event VoteCommitted(address indexed voter, bytes32 commitment, uint256 timestamp);

    /// @notice Emitted when a vote is revealed
    event VoteRevealed(address indexed voter, uint256 choice, uint256 timestamp);

    /// @notice Emitted when poll is closed
    event PollClosed(uint256 timestamp);

    /// @notice Emitted when results are tallied
    event ResultsTallied(uint256[] results, uint256 timestamp);

    /// @notice Commit a vote with ZK proof
    function commitVote(bytes32 commitment, bytes calldata zkProof, bytes32[] calldata merklePath) external;

    /// @notice Reveal a committed vote
    function revealVote(uint256 choice, bytes32 salt) external;

    /// @notice Close the poll (called by oracle)
    function closePoll() external;

    /// @notice Tally the final results
    function tally() external returns (uint256[] memory);

    /// @notice Get current poll state
    function state() external view returns (PollState);

    /// @notice Get poll question
    function question() external view returns (string memory);

    /// @notice Get poll options
    function options() external view returns (string[] memory);

    /// @notice Get voter's commitment
    function getCommitment(address voter) external view returns (VoteCommitment memory);

    /// @notice Get tallied results
    function getResults() external view returns (uint256[] memory);
}
