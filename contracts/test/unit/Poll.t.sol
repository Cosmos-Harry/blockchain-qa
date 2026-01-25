// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../../src/Poll.sol";
import "../../src/MockOracle.sol";
import "../../src/MockZKVerifier.sol";

/**
 * @title PollTest
 * @notice Unit tests for Poll contract
 */
contract PollTest is Test {
    Poll public poll;
    MockOracle public oracle;
    MockZKVerifier public verifier;

    address public creator = address(1);
    address public voter1 = address(2);
    address public voter2 = address(3);
    address public attacker = address(4);

    string public question = "Favorite color?";
    string[] public options;
    uint256 public duration = 1 hours;
    bytes32 public voterMerkleRoot;

    // Merkle tree for 3 voters (voter1, voter2, and one more)
    bytes32[] public merkleProof1;
    bytes32[] public merkleProof2;

    function setUp() public {
        // Setup options
        options.push("Red");
        options.push("Blue");
        options.push("Green");

        // Deploy mock contracts
        verifier = new MockZKVerifier(MockZKVerifier.VerificationMode.AlwaysPass);
        oracle = new MockOracle(MockOracle.ResponseMode.OnTime, 0);

        // Create simple Merkle tree for testing
        // Tree: root
        //       /    \
        //    h1       h2
        //   /  \     /
        // v1   v2  v3

        bytes32 leaf1 = keccak256(abi.encodePacked(voter1));
        bytes32 leaf2 = keccak256(abi.encodePacked(voter2));
        bytes32 leaf3 = keccak256(abi.encodePacked(address(5)));

        bytes32 h1 = keccak256(abi.encodePacked(leaf1, leaf2));
        bytes32 h2 = keccak256(abi.encodePacked(leaf3));

        voterMerkleRoot = keccak256(abi.encodePacked(h1, h2));

        // Merkle proofs
        merkleProof1.push(leaf2);
        merkleProof1.push(h2);

        merkleProof2.push(leaf1);
        merkleProof2.push(h2);

        // Deploy poll
        vm.prank(creator);
        poll = new Poll(
            question,
            options,
            duration,
            voterMerkleRoot,
            address(verifier),
            address(oracle)
        );
    }

    // ========== Constructor Tests ==========

    function testConstructor() public {
        assertEq(poll.question(), question);
        assertEq(poll.creator(), creator);
        assertEq(poll.duration(), duration);
        assertEq(poll.voterMerkleRoot(), voterMerkleRoot);
        assertEq(address(poll.zkVerifier()), address(verifier));
        assertEq(address(poll.oracle()), address(oracle));
        assertEq(uint256(poll.state()), uint256(IPoll.PollState.Active));
    }

    function testConstructorRequiresTwoOptions() public {
        string[] memory badOptions = new string[](1);
        badOptions[0] = "Only one";

        vm.expectRevert();
        new Poll(
            question,
            badOptions,
            duration,
            voterMerkleRoot,
            address(verifier),
            address(oracle)
        );
    }

    // ========== Vote Commitment Tests ==========

    function testCommitVote() public {
        bytes32 commitment = keccak256(abi.encodePacked(uint256(0), bytes32(uint256(123)), voter1));
        bytes memory proof = hex"dead";

        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);

        IPoll.VoteCommitment memory vc = poll.getCommitment(voter1);
        assertEq(vc.commitment, commitment);
        assertEq(vc.revealed, false);
        assertEq(poll.totalCommitted(), 1);
    }

    function testCannotCommitTwice() public {
        bytes32 commitment = keccak256(abi.encodePacked(uint256(0), bytes32(uint256(123)), voter1));
        bytes memory proof = hex"dead";

        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);

        // Try to commit again
        vm.expectRevert(Poll.AlreadyVoted.selector);
        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);
    }

    function testCannotCommitAfterClose() public {
        // Fast forward past end time
        vm.warp(poll.endTime() + 1);

        // Oracle closes poll
        oracle.fulfillRequest(address(poll));

        assertEq(uint256(poll.state()), uint256(IPoll.PollState.Closed));

        // Try to commit
        bytes32 commitment = keccak256(abi.encodePacked(uint256(0), bytes32(uint256(123)), voter1));
        bytes memory proof = hex"dead";

        vm.expectRevert(Poll.PollNotActive.selector);
        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);
    }

    function testCannotCommitWithInvalidMerkleProof() public {
        bytes32 commitment = keccak256(abi.encodePacked(uint256(0), bytes32(uint256(123)), attacker));
        bytes memory proof = hex"dead";

        bytes32[] memory badProof = new bytes32[](2);
        badProof[0] = bytes32(uint256(999));
        badProof[1] = bytes32(uint256(888));

        vm.expectRevert(Poll.InvalidMerkleProof.selector);
        vm.prank(attacker);
        poll.commitVote(commitment, proof, badProof);
    }

    function testCannotCommitWithInvalidZKProof() public {
        // Set verifier to always fail
        verifier.setMode(MockZKVerifier.VerificationMode.AlwaysFail);

        bytes32 commitment = keccak256(abi.encodePacked(uint256(0), bytes32(uint256(123)), voter1));
        bytes memory proof = hex"bad";

        vm.expectRevert(Poll.InvalidProof.selector);
        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);
    }

    // ========== Vote Reveal Tests ==========

    function testRevealVote() public {
        // Commit
        uint256 choice = 1;
        bytes32 salt = bytes32(uint256(123));
        bytes32 commitment = keccak256(abi.encodePacked(choice, salt, voter1));
        bytes memory proof = hex"dead";

        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);

        // Close poll
        vm.warp(poll.endTime() + 1);
        oracle.fulfillRequest(address(poll));

        // Reveal
        vm.prank(voter1);
        poll.revealVote(choice, salt);

        IPoll.VoteCommitment memory vc = poll.getCommitment(voter1);
        assertEq(vc.revealed, true);
        assertEq(poll.totalRevealed(), 1);
    }

    function testCannotRevealBeforeClose() public {
        // Commit
        uint256 choice = 1;
        bytes32 salt = bytes32(uint256(123));
        bytes32 commitment = keccak256(abi.encodePacked(choice, salt, voter1));
        bytes memory proof = hex"dead";

        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);

        // Try to reveal without closing
        vm.expectRevert(Poll.PollNotClosed.selector);
        vm.prank(voter1);
        poll.revealVote(choice, salt);
    }

    function testCannotRevealWithoutCommitment() public {
        // Close poll without any commitments
        vm.warp(poll.endTime() + 1);
        oracle.fulfillRequest(address(poll));

        vm.expectRevert(Poll.NoCommitment.selector);
        vm.prank(voter1);
        poll.revealVote(0, bytes32(0));
    }

    function testCannotRevealTwice() public {
        // Commit and close
        uint256 choice = 1;
        bytes32 salt = bytes32(uint256(123));
        bytes32 commitment = keccak256(abi.encodePacked(choice, salt, voter1));
        bytes memory proof = hex"dead";

        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);

        vm.warp(poll.endTime() + 1);
        oracle.fulfillRequest(address(poll));

        // First reveal
        vm.prank(voter1);
        poll.revealVote(choice, salt);

        // Second reveal
        vm.expectRevert(Poll.AlreadyRevealed.selector);
        vm.prank(voter1);
        poll.revealVote(choice, salt);
    }

    function testCannotRevealWithWrongData() public {
        // Commit
        uint256 choice = 1;
        bytes32 salt = bytes32(uint256(123));
        bytes32 commitment = keccak256(abi.encodePacked(choice, salt, voter1));
        bytes memory proof = hex"dead";

        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);

        vm.warp(poll.endTime() + 1);
        oracle.fulfillRequest(address(poll));

        // Reveal with wrong data
        vm.expectRevert(Poll.InvalidReveal.selector);
        vm.prank(voter1);
        poll.revealVote(2, salt); // Wrong choice
    }

    function testCannotRevealInvalidChoice() public {
        // Commit
        uint256 choice = 0;
        bytes32 salt = bytes32(uint256(123));
        bytes32 commitment = keccak256(abi.encodePacked(choice, salt, voter1));
        bytes memory proof = hex"dead";

        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);

        vm.warp(poll.endTime() + 1);
        oracle.fulfillRequest(address(poll));

        // Try to reveal with out-of-range choice
        vm.expectRevert(Poll.InvalidChoice.selector);
        vm.prank(voter1);
        poll.revealVote(999, salt);
    }

    // ========== Poll Closing Tests ==========

    function testClosePoll() public {
        vm.warp(poll.endTime() + 1);

        vm.prank(address(oracle));
        poll.closePoll();

        assertEq(uint256(poll.state()), uint256(IPoll.PollState.Closed));
    }

    function testOnlyOracleCanClose() public {
        vm.warp(poll.endTime() + 1);

        vm.expectRevert(Poll.UnauthorizedOracle.selector);
        vm.prank(attacker);
        poll.closePoll();
    }

    function testCannotCloseAlreadyClosedPoll() public {
        vm.warp(poll.endTime() + 1);

        vm.prank(address(oracle));
        poll.closePoll();

        vm.expectRevert(Poll.PollAlreadyClosed.selector);
        vm.prank(address(oracle));
        poll.closePoll();
    }

    // ========== Tally Tests ==========

    function testTally() public {
        // Two voters commit different choices
        uint256 choice1 = 0;
        bytes32 salt1 = bytes32(uint256(123));
        bytes32 commitment1 = keccak256(abi.encodePacked(choice1, salt1, voter1));

        uint256 choice2 = 1;
        bytes32 salt2 = bytes32(uint256(456));
        bytes32 commitment2 = keccak256(abi.encodePacked(choice2, salt2, voter2));

        bytes memory proof = hex"dead";

        vm.prank(voter1);
        poll.commitVote(commitment1, proof, merkleProof1);

        vm.prank(voter2);
        poll.commitVote(commitment2, proof, merkleProof2);

        // Close and reveal
        vm.warp(poll.endTime() + 1);
        oracle.fulfillRequest(address(poll));

        vm.prank(voter1);
        poll.revealVote(choice1, salt1);

        vm.prank(voter2);
        poll.revealVote(choice2, salt2);

        // Tally
        uint256[] memory results = poll.tally();

        assertEq(results[0], 1); // One vote for option 0
        assertEq(results[1], 1); // One vote for option 1
        assertEq(results[2], 0); // Zero votes for option 2
        assertEq(uint256(poll.state()), uint256(IPoll.PollState.Tallied));
    }

    function testCannotTallyBeforeClose() public {
        vm.expectRevert(Poll.PollNotClosed.selector);
        poll.tally();
    }

    function testTallyWithZeroVoters() public {
        // Close poll without any votes
        vm.warp(poll.endTime() + 1);
        oracle.fulfillRequest(address(poll));

        uint256[] memory results = poll.tally();

        assertEq(results[0], 0);
        assertEq(results[1], 0);
        assertEq(results[2], 0);
    }

    function testGetResultsOnlyAfterTally() public {
        // Close poll
        vm.warp(poll.endTime() + 1);
        oracle.fulfillRequest(address(poll));

        // Cannot get results before tally
        vm.expectRevert(Poll.PollNotTallied.selector);
        poll.getResults();

        // Tally
        poll.tally();

        // Now can get results
        uint256[] memory results = poll.getResults();
        assertEq(results.length, 3);
    }

    // ========== Events Tests ==========

    function testVoteCommittedEvent() public {
        bytes32 commitment = keccak256(abi.encodePacked(uint256(0), bytes32(uint256(123)), voter1));
        bytes memory proof = hex"dead";

        vm.expectEmit(true, false, false, true);
        emit IPoll.VoteCommitted(voter1, commitment, block.timestamp);

        vm.prank(voter1);
        poll.commitVote(commitment, proof, merkleProof1);
    }

    function testVoteRevealedEvent() public {
        // Setup: commit and close
        uint256 choice = 1;
        bytes32 salt = bytes32(uint256(123));
        bytes32 commitment = keccak256(abi.encodePacked(choice, salt, voter1));

        vm.prank(voter1);
        poll.commitVote(commitment, hex"dead", merkleProof1);

        vm.warp(poll.endTime() + 1);
        oracle.fulfillRequest(address(poll));

        // Expect reveal event
        vm.expectEmit(true, false, false, true);
        emit IPoll.VoteRevealed(voter1, choice, block.timestamp);

        vm.prank(voter1);
        poll.revealVote(choice, salt);
    }
}
