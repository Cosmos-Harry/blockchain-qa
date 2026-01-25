// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IOracle} from "./interfaces/IOracle.sol";
import {IPoll} from "./interfaces/IPoll.sol";

/**
 * @title MockOracle
 * @notice Configurable oracle for testing poll closing scenarios
 * @dev Supports multiple response modes:
 *      - OnTime: Closes poll exactly at scheduled time
 *      - Late: Closes poll with configured delay
 *      - Invalid: Sends invalid data
 *      - NoResponse: Doesn't respond (manual triggering required)
 */
contract MockOracle is IOracle {
    /// @notice Response mode enum
    enum ResponseMode {
        OnTime,
        Late,
        Invalid,
        NoResponse
    }

    /// @notice Current response mode
    ResponseMode public mode;

    /// @notice Delay in seconds for Late mode
    uint256 public lateDelay;

    /// @notice Owner who can change modes
    address public immutable owner;

    /// @notice Mapping of poll address to end time
    mapping(address => uint256) public pollEndTimes;

    /// @notice Mapping to track fulfilled requests
    mapping(address => bool) public fulfilled;

    /// @notice Custom errors
    error UnauthorizedCaller();
    error PollNotRegistered();
    error TooEarlyToClose();
    error RequestAlreadyFulfilled();

    /// @notice Emitted when mode is changed
    event ModeChanged(ResponseMode newMode, uint256 delay);

    /**
     * @notice Constructor
     * @param _mode Initial response mode
     * @param _lateDelay Delay for Late mode (in seconds)
     */
    constructor(ResponseMode _mode, uint256 _lateDelay) {
        owner = msg.sender;
        mode = _mode;
        lateDelay = _lateDelay;
    }

    /**
     * @notice Set response mode (owner only)
     * @param _mode New response mode
     * @param _delay New delay for Late mode
     */
    function setMode(ResponseMode _mode, uint256 _delay) external {
        if (msg.sender != owner) revert UnauthorizedCaller();

        mode = _mode;
        lateDelay = _delay;

        emit ModeChanged(_mode, _delay);
    }

    /**
     * @notice Request poll closing at specified time
     * @param poll Address of the poll contract
     * @param endTime Unix timestamp when poll should close
     */
    function requestPollClose(address poll, uint256 endTime) external override {
        pollEndTimes[poll] = endTime;
        fulfilled[poll] = false;

        emit PollCloseRequested(poll, endTime);

        // In OnTime mode, we would typically trigger this automatically
        // For testing, we'll use a manual trigger via fulfillRequest()
    }

    /**
     * @notice Fulfill a poll close request
     * @param poll Address of the poll contract
     */
    function fulfillRequest(address poll) external override {
        if (pollEndTimes[poll] == 0) revert PollNotRegistered();
        if (fulfilled[poll]) revert RequestAlreadyFulfilled();

        uint256 endTime = pollEndTimes[poll];
        uint256 requiredTime = endTime;

        // Apply mode-specific behavior
        if (mode == ResponseMode.OnTime) {
            // Close exactly at end time
            if (block.timestamp < endTime) revert TooEarlyToClose();
        } else if (mode == ResponseMode.Late) {
            // Close after delay
            requiredTime = endTime + lateDelay;
            if (block.timestamp < requiredTime) revert TooEarlyToClose();
        } else if (mode == ResponseMode.Invalid) {
            // Send invalid data (try to close before end time)
            // This should fail at the Poll contract level
            // No time check here to allow testing invalid behavior
        } else if (mode == ResponseMode.NoResponse) {
            // In NoResponse mode, manual trigger is required
            // Anyone can call this function to manually close the poll
        }

        // Mark as fulfilled
        fulfilled[poll] = true;

        // Call poll's closePoll function
        IPoll(poll).closePoll();

        emit PollCloseFulfilled(poll, block.timestamp);
    }

    /**
     * @notice Manually fulfill request (for testing)
     * @param poll Address of the poll
     */
    function manualFulfill(address poll) external {
        if (msg.sender != owner) revert UnauthorizedCaller();

        fulfilled[poll] = true;
        IPoll(poll).closePoll();

        emit PollCloseFulfilled(poll, block.timestamp);
    }

    /**
     * @notice Check if a poll can be closed now
     * @param poll Address of the poll
     */
    function canClose(address poll) external view returns (bool) {
        if (pollEndTimes[poll] == 0 || fulfilled[poll]) {
            return false;
        }

        uint256 endTime = pollEndTimes[poll];

        if (mode == ResponseMode.OnTime) {
            return block.timestamp >= endTime;
        } else if (mode == ResponseMode.Late) {
            return block.timestamp >= endTime + lateDelay;
        } else if (mode == ResponseMode.Invalid || mode == ResponseMode.NoResponse) {
            return true; // Can attempt anytime (may fail)
        }

        return false;
    }
}
