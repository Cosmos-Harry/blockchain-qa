// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IZKVerifier
 * @notice Interface for Groth16 ZK proof verifier
 */
interface IZKVerifier {
    /**
     * @notice Verify a Groth16 proof
     * @param proof The proof data (a, b, c points)
     * @param publicInputs The public inputs to the circuit
     * @return True if proof is valid
     */
    function verifyProof(bytes calldata proof, uint256[] calldata publicInputs) external view returns (bool);
}
