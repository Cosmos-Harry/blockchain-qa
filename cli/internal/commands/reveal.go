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
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	revealCmd.Flags().StringVar(&pollAddress, "poll", "", "Poll contract address (or reads from state)")
	revealCmd.Flags().Uint64Var(&revealChoice, "choice", 0, "Vote choice (or reads from state)")
	revealCmd.Flags().StringVar(&revealNonce, "nonce", "", "Nonce used in commitment (or reads from state)")
}

func runReveal(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	// Create wallet
	w, err := wallet.NewWallet(ctx, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	defer w.Close()

	// Resolve poll, nonce, and choice from flags or saved state
	state := loadState()
	resolved, err := resolveFlag(pollAddress, state.PollAddress, "poll")
	if err != nil {
		return err
	}
	pollAddress = resolved

	resolved, err = resolveFlag(revealNonce, state.Nonce, "nonce")
	if err != nil {
		return err
	}
	revealNonce = resolved

	if !cmd.Flags().Changed("choice") && state.Choice > 0 {
		revealChoice = state.Choice
	}

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
		revertReason := decodeRevealRevertReason(ctx, w, pollAddr, auth, big.NewInt(int64(revealChoice)), salt)
		return fmt.Errorf("transaction reverted: %s", revertReason)
	}

	log.Printf("\nVote revealed successfully!\n")
	log.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
	log.Printf("Block Number: %d\n", receipt.BlockNumber.Uint64())
	log.Printf("Gas Used: %d\n", receipt.GasUsed)

	return nil
}

func decodeRevealRevertReason(ctx context.Context, w *wallet.Wallet, pollAddr common.Address, auth *bind.TransactOpts, choice *big.Int, salt [32]byte) string {
	pollABI, err := abi.JSON(strings.NewReader(bindings.PollABI))
	if err != nil {
		return "unknown (failed to parse ABI)"
	}

	data, err := pollABI.Pack("revealVote", choice, salt)
	if err != nil {
		return "unknown (failed to pack call data)"
	}

	msg := ethereum.CallMsg{
		From: auth.From,
		To:   &pollAddr,
		Data: data,
	}

	_, err = w.GetClient().CallContract(ctx, msg, nil)
	if err != nil {
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
