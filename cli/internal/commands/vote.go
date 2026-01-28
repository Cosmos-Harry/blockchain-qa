package commands

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/bindings"
	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/wallet"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var (
	pollAddress  string
	choice       uint64
	merkleProof  string
)

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Cast a vote in a poll",
	Long:  `Submit a vote commitment with zero-knowledge proof to a poll.`,
	RunE:  runVote,
}

func init() {
	rootCmd.AddCommand(voteCmd)

	voteCmd.Flags().StringVar(&pollAddress, "poll", "", "Poll contract address (or reads from state)")
	voteCmd.Flags().Uint64Var(&choice, "choice", 0, "Vote choice (0-indexed)")
	voteCmd.Flags().StringVar(&merkleProof, "proof", "", "Merkle proof for voter eligibility (comma-separated hashes)")
}

func runVote(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	// Create wallet
	w, err := wallet.NewWallet(ctx, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	defer w.Close()

	// Resolve poll address from flag or saved state
	state := loadState()
	resolved, err := resolveFlag(pollAddress, state.PollAddress, "poll")
	if err != nil {
		return err
	}
	pollAddress = resolved

	log.Printf("Voting from address: %s\n", w.Address().Hex())
	log.Printf("Poll: %s\n", pollAddress)
	log.Printf("Choice: %d\n", choice)

	// Generate random nonce for privacy
	nonce := make([]byte, 32)
	if _, err := rand.Read(nonce); err != nil {
		return fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Compute commitment: Hash(choice || nonce || voter)
	commitment := computeCommitment(choice, nonce, w.Address())
	log.Printf("Commitment: %s\n", commitment.Hex())

	// TODO: Generate ZK proof using Rust prover (via FFI or subprocess)
	// For now, create a dummy proof
	log.Println("\nGenerating zero-knowledge proof...")
	zkProof := generateDummyZKProof()
	log.Printf("ZK Proof: %s\n", hex.EncodeToString(zkProof))

	// Create Poll contract instance
	log.Println("\nSubmitting vote to contract...")
	pollAddr := common.HexToAddress(pollAddress)
	poll, err := bindings.NewPoll(pollAddr, w.GetClient())
	if err != nil {
		return fmt.Errorf("failed to create poll instance: %w", err)
	}

	// Get auth
	auth, err := w.GetAuth(ctx)
	if err != nil {
		return fmt.Errorf("failed to get auth: %w", err)
	}

	log.Printf("  Gas Limit: %d\n", auth.GasLimit)
	log.Printf("  Gas Price: %s wei\n", auth.GasPrice.String())

	// Parse merkle proof if provided, or use empty array for single-voter trees
	var proofHashes [][32]byte
	if merkleProof != "" {
		// Parse comma-separated hex hashes
		proofStrs := strings.Split(merkleProof, ",")
		proofHashes = make([][32]byte, len(proofStrs))
		for i, ps := range proofStrs {
			ps = strings.TrimSpace(ps)
			hashBytes := common.FromHex(ps)
			if len(hashBytes) != 32 {
				return fmt.Errorf("invalid proof hash at index %d: expected 32 bytes, got %d", i, len(hashBytes))
			}
			copy(proofHashes[i][:], hashBytes)
		}
	}
	// If no proof provided, use empty array (works for single-voter merkle trees)
	log.Printf("Merkle proof hashes count: %d\n", len(proofHashes))
	log.Printf("Commitment: %s\n", commitment.Hex())
	log.Printf("ZK Proof length: %d bytes\n", len(zkProof))

	// Call commitVote on the contract
	tx, err := poll.CommitVote(auth, commitment, zkProof, proofHashes)
	if err != nil {
		return fmt.Errorf("failed to commit vote: %w", err)
	}

	log.Printf("\nTransaction submitted: %s\n", tx.Hash().Hex())
	log.Println("Waiting for confirmation...")

	// Wait for transaction receipt
	receipt, err := w.WaitForReceipt(ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("failed to get receipt: %w", err)
	}

	if receipt.Status == 0 {
		// Decode revert reason from receipt logs or re-call to get error
		revertReason := decodeRevertReason(ctx, w, pollAddr, auth, commitment, zkProof, proofHashes)
		return fmt.Errorf("transaction reverted: %s", revertReason)
	}

	// Save nonce and choice to state for reveal step
	state.Nonce = hex.EncodeToString(nonce)
	state.Choice = choice
	if err := saveState(state); err != nil {
		log.Printf("Warning: failed to save state: %v\n", err)
	}

	log.Printf("\nVote committed successfully!\n")
	log.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
	log.Printf("Block Number: %d\n", receipt.BlockNumber.Uint64())
	log.Printf("Gas Used: %d\n", receipt.GasUsed)
	log.Printf("Nonce saved to state (needed for reveal)\n")

	return nil
}

// computeCommitment computes the vote commitment hash matching the contract's
// keccak256(abi.encodePacked(uint256 choice, bytes32 salt, address voter))
func computeCommitment(choice uint64, nonce []byte, voter common.Address) common.Hash {
	// uint256: left-pad choice to 32 bytes (big-endian)
	choiceBig := new(big.Int).SetUint64(choice)
	choiceBytes := common.LeftPadBytes(choiceBig.Bytes(), 32)

	// bytes32: nonce is already 32 bytes
	// address: 20 bytes, no padding in encodePacked
	var data []byte
	data = append(data, choiceBytes...)
	data = append(data, nonce...)
	data = append(data, voter.Bytes()...)
	return crypto.Keccak256Hash(data)
}

// generateDummyZKProof generates a dummy ZK proof for demonstration
func generateDummyZKProof() []byte {
	// In production, this would call the Rust ZK prover via FFI or subprocess
	// The prover would generate a real Groth16 proof
	proof := make([]byte, 192) // Groth16 proofs are ~192 bytes
	if _, err := rand.Read(proof); err != nil {
		log.Fatalf("Failed to generate random proof: %v", err)
	}
	return proof
}

// decodeRevertReason re-calls the contract via eth_call to extract the revert reason
func decodeRevertReason(ctx context.Context, w *wallet.Wallet, pollAddr common.Address, auth *bind.TransactOpts, commitment common.Hash, zkProof []byte, proofHashes [][32]byte) string {
	// Parse the Poll ABI
	pollABI, err := abi.JSON(strings.NewReader(bindings.PollABI))
	if err != nil {
		return "unknown (failed to parse ABI)"
	}

	// Pack the commitVote call data
	data, err := pollABI.Pack("commitVote", commitment, zkProof, proofHashes)
	if err != nil {
		return "unknown (failed to pack call data)"
	}

	// Execute as eth_call to get revert reason
	msg := ethereum.CallMsg{
		From: auth.From,
		To:   &pollAddr,
		Data: data,
	}

	_, err = w.GetClient().CallContract(ctx, msg, nil)
	if err != nil {
		// Try to decode custom error selector into error name
		errStr := err.Error()
		for _, e := range pollABI.Errors {
			if strings.Contains(errStr, fmt.Sprintf("0x%x", e.ID[:4])) {
				return e.Name
			}
		}
		return errStr
	}
	return "unknown"
}
