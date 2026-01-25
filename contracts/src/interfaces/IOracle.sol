// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IOracle
 * @notice Interface for oracle that triggers poll closing
 */
interface IOracle {
    /// @notice Emitted when a poll close request is made
    event PollCloseRequested(address indexed poll, uint256 endTime);

    /// @notice Emitted when a poll close request is fulfilled
    event PollCloseFulfilled(address indexed poll, uint256 timestamp);

    /**
     * @notice Request poll closing at specified time
     * @param poll Address of the poll contract
     * @param endTime Unix timestamp when poll should close
     */
    function requestPollClose(address poll, uint256 endTime) external;

    /**
     * @notice Fulfill a poll close request
     * @param poll Address of the poll contract
     */
    function fulfillRequest(address poll) external;
}
