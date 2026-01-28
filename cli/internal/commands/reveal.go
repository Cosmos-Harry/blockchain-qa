package commands

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/bindings"
	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/wallet"
	"github.com/ethereum/go-ethereum/common"
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

	mustMarkRequired(revealCmd, "poll")
	mustMarkRequired(revealCmd, "choice")
	mustMarkRequired(revealCmd, "nonce")
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

	// Parse nonce (strip 0x prefix if present)
	nonceHex := strings.TrimPrefix(revealNonce, "0x")
	nonceBytes, err := hex.DecodeString(nonceHex)
	if err != nil {
		return fmt.Errorf("failed to decode nonce: %w", err)
	}

	if len(nonceBytes) != 32 {
		return fmt.Errorf("nonce must be 32 bytes, got %d", len(nonceBytes))
	}

	// Verify commitment matches (locally)
	commitment := computeCommitment(revealChoice, nonceBytes, w.Address())
	log.Printf("Recomputed Commitment: %s\n", commitment.Hex())

	// Create Poll contract instance
	log.Println("\nSubmitting reveal to contract...")
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

	// Convert nonce to [32]byte
	var salt [32]byte
	copy(salt[:], nonceBytes)

	// Call revealVote on the contract
	tx, err := poll.RevealVote(auth, big.NewInt(int64(revealChoice)), salt)
	if err != nil {
		return fmt.Errorf("failed to reveal vote: %w", err)
	}

	log.Printf("\nTransaction submitted: %s\n", tx.Hash().Hex())
	log.Println("Waiting for confirmation...")

	// Wait for transaction receipt
	receipt, err := w.WaitForReceipt(ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("failed to get receipt: %w", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	log.Printf("\nVote revealed successfully!\n")
	log.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
	log.Printf("Block Number: %d\n", receipt.BlockNumber.Uint64())
	log.Printf("Gas Used: %d\n", receipt.GasUsed)

	return nil
}
