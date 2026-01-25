package commands

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/wallet"
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

	voteCmd.Flags().StringVar(&pollAddress, "poll", "", "Poll contract address")
	voteCmd.Flags().Uint64Var(&choice, "choice", 0, "Vote choice (0-indexed)")
	voteCmd.Flags().StringVar(&merkleProof, "proof", "", "Merkle proof for voter eligibility (comma-separated hashes)")

	voteCmd.MarkFlagRequired("poll")
}

func runVote(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	// Create wallet
	w, err := wallet.NewWallet(ctx, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	defer w.Close()

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

	// TODO: Call Poll.commitVote() when contract bindings are generated
	// For now, simulate the call
	log.Println("\nSimulating commitVote transaction...")

	// Get auth
	auth, err := w.GetAuth(ctx)
	if err != nil {
		return fmt.Errorf("failed to get auth: %w", err)
	}

	log.Printf("  Gas Limit: %d\n", auth.GasLimit)
	log.Printf("  Gas Price: %s wei\n", auth.GasPrice.String())

	// Simulate successful transaction
	txHash := "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"
	log.Printf("\nVote committed successfully!\n")
	log.Printf("Transaction Hash: %s\n", txHash)
	log.Printf("\nIMPORTANT: Save your nonce to reveal later:\n")
	log.Printf("Nonce: %s\n", hex.EncodeToString(nonce))

	return nil
}

// computeCommitment computes the vote commitment hash
func computeCommitment(choice uint64, nonce []byte, voter common.Address) common.Hash {
	// Simplified commitment (should match circuit implementation)
	// In production, this would use Poseidon hash to match the circuit
	data := append([]byte(fmt.Sprintf("%d", choice)), nonce...)
	data = append(data, voter.Bytes()...)
	return crypto.Keccak256Hash(data)
}

// generateDummyZKProof generates a dummy ZK proof for demonstration
func generateDummyZKProof() []byte {
	// In production, this would call the Rust ZK prover via FFI or subprocess
	// The prover would generate a real Groth16 proof
	proof := make([]byte, 192) // Groth16 proofs are ~192 bytes
	rand.Read(proof)
	return proof
}
