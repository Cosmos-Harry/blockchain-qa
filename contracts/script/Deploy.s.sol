// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/PollFactory.sol";
import "../src/MockOracle.sol";
import "../src/MockZKVerifier.sol";

/**
 * @title Deploy
 * @notice Deployment script for local development
 * @dev Deploys all contracts needed for the poll system
 */
contract Deploy is Script {
    function run() external {
        // Get private key from environment or use default Anvil key
        uint256 deployerPrivateKey =
            vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));

        vm.startBroadcast(deployerPrivateKey);

        // Deploy MockZKVerifier (always passes for local testing)
        MockZKVerifier zkVerifier = new MockZKVerifier(MockZKVerifier.VerificationMode.AlwaysPass);
        console.log("MockZKVerifier deployed to:", address(zkVerifier));

        // Deploy MockOracle (OnTime mode, no delay)
        MockOracle oracle = new MockOracle(
            MockOracle.ResponseMode.OnTime,
            0 // no delay
        );
        console.log("MockOracle deployed to:", address(oracle));

        // Deploy PollFactory
        PollFactory factory = new PollFactory(address(zkVerifier), address(oracle));
        console.log("PollFactory deployed to:", address(factory));

        vm.stopBroadcast();

        // Print deployment summary
        console.log("\n=== Deployment Summary ===");
        console.log("ZK Verifier:", address(zkVerifier));
        console.log("Oracle:     ", address(oracle));
        console.log("Factory:    ", address(factory));
        console.log("\nSave these addresses for CLI usage!");
    }
}
