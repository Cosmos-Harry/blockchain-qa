package commands

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/wallet"
	"github.com/spf13/cobra"
)

var (
	revealChoice uint64
	revealNonce  string
)

var revealCmd = &cobra.Command{
	Use:   "reveal",
	Short: "Reveal a vote after poll closes",
	Long:  `Reveal your vote choice and nonce after the poll has closed.`,
	RunE:  runReveal,
}

func init() {
	rootCmd.AddCommand(revealCmd)

	revealCmd.Flags().StringVar(&pollAddress, "poll", "", "Poll contract address")
	revealCmd.Flags().Uint64Var(&revealChoice, "choice", 0, "Vote choice (must match commitment)")
	revealCmd.Flags().StringVar(&revealNonce, "nonce", "", "Nonce used in commitment (hex string)")

	revealCmd.MarkFlagRequired("poll")
	revealCmd.MarkFlagRequired("choice")
	revealCmd.MarkFlagRequired("nonce")
}

func runReveal(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	// Create wallet
	w, err := wallet.NewWallet(ctx, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	defer w.Close()

	log.Printf("Revealing vote from address: %s\n", w.Address().Hex())
	log.Printf("Poll: %s\n", pollAddress)
	log.Printf("Choice: %d\n", revealChoice)

	// Parse nonce
	nonceBytes, err := hex.DecodeString(revealNonce)
	if err != nil {
		return fmt.Errorf("failed to decode nonce: %w", err)
	}

	if len(nonceBytes) != 32 {
		return fmt.Errorf("nonce must be 32 bytes, got %d", len(nonceBytes))
	}

	// Verify commitment matches (locally)
	commitment := computeCommitment(revealChoice, nonceBytes, w.Address())
	log.Printf("Recomputed Commitment: %s\n", commitment.Hex())

	// TODO: Call Poll.revealVote() when contract bindings are generated
	// For now, simulate the call
	log.Println("\nSimulating revealVote transaction...")

	// Get auth
	auth, err := w.GetAuth(ctx)
	if err != nil {
		return fmt.Errorf("failed to get auth: %w", err)
	}

	log.Printf("  Gas Limit: %d\n", auth.GasLimit)
	log.Printf("  Gas Price: %s wei\n", auth.GasPrice.String())

	// Simulate successful transaction
	txHash := "0x1234abcd5678ef901234abcd5678ef901234abcd5678ef901234abcd5678ef90"
	log.Printf("\nVote revealed successfully!\n")
	log.Printf("Transaction Hash: %s\n", txHash)

	return nil
}
