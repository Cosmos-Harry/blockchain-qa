// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IZKVerifier} from "./interfaces/IZKVerifier.sol";

/**
 * @title MockZKVerifier
 * @notice Mock ZK verifier for testing purposes
 * @dev In production, this will be replaced with a real Groth16 verifier
 *      generated from the Rust circuit using snarkjs
 */
contract MockZKVerifier is IZKVerifier {
    /// @notice Mode for controlling verification behavior
    enum VerificationMode {
        AlwaysPass,      // Always returns true
        AlwaysFail,      // Always returns false
        CheckPublicInput // Validates public input format
    }

    /// @notice Current verification mode
    VerificationMode public mode;

    /// @notice Owner who can change mode
    address public immutable owner;

    /// @notice Number of verification calls
    uint256 public verificationCount;

    /// @notice Custom errors
    error UnauthorizedCaller();

    /// @notice Emitted when verification is attempted
    event VerificationAttempted(bool result, uint256 publicInputCount);

    /// @notice Emitted when mode is changed
    event ModeChanged(VerificationMode newMode);

    /**
     * @notice Constructor
     * @param _mode Initial verification mode
     */
    constructor(VerificationMode _mode) {
        owner = msg.sender;
        mode = _mode;
    }

    /**
     * @notice Set verification mode (owner only)
     * @param _mode New verification mode
     */
    function setMode(VerificationMode _mode) external {
        if (msg.sender != owner) revert UnauthorizedCaller();
        mode = _mode;
        emit ModeChanged(_mode);
    }

    /**
     * @notice Verify a ZK proof (mock implementation)
     * @param proof The proof data (ignored in mock)
     * @param publicInputs The public inputs to verify
     * @return True if proof is valid according to current mode
     */
    function verifyProof(
        bytes calldata proof,
        uint256[] calldata publicInputs
    ) external override returns (bool) {
        verificationCount++;

        bool result;

        if (mode == VerificationMode.AlwaysPass) {
            result = true;
        } else if (mode == VerificationMode.AlwaysFail) {
            result = false;
        } else if (mode == VerificationMode.CheckPublicInput) {
            // Validate that we have at least one public input (commitment)
            result = publicInputs.length > 0 && publicInputs[0] != 0;
        }

        emit VerificationAttempted(result, publicInputs.length);

        return result;
    }

    /**
     * @notice Get verification statistics
     * @return Total number of verification attempts
     */
    function getVerificationCount() external view returns (uint256) {
        return verificationCount;
    }
}
