package commands

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/bindings"
	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	question         string
	options          string
	duration         int64
	voterMerkleRoot  string
	pollFactoryAddr  string
)

var createPollCmd = &cobra.Command{
	Use:   "create-poll",
	Short: "Create a new poll",
	Long:  `Create a new confidential poll with the specified question and options.`,
	RunE:  runCreatePoll,
}

func init() {
	rootCmd.AddCommand(createPollCmd)

	createPollCmd.Flags().StringVar(&question, "question", "", "Poll question")
	createPollCmd.Flags().StringVar(&options, "options", "", "Comma-separated voting options")
	createPollCmd.Flags().Int64Var(&duration, "duration", 3600, "Poll duration in seconds")
	createPollCmd.Flags().StringVar(&voterMerkleRoot, "voter-root", "", "Merkle root of eligible voters (defaults to zero root = anyone can vote)")
	createPollCmd.Flags().StringVar(&pollFactoryAddr, "factory", "", "PollFactory contract address (uses POLL_FACTORY_ADDRESS env var if not specified)")

	mustMarkRequired(createPollCmd, "question")
	mustMarkRequired(createPollCmd, "options")
}

func runCreatePoll(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	// Create wallet
	w, err := wallet.NewWallet(ctx, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	defer w.Close()

	log.Printf("Creating poll from address: %s\n", w.Address().Hex())

	// Parse options
	optionsList := strings.Split(options, ",")
	for i := range optionsList {
		optionsList[i] = strings.TrimSpace(optionsList[i])
	}

	if len(optionsList) < 2 {
		return fmt.Errorf("at least 2 options are required")
	}

	// Get factory address from flag or environment
	if pollFactoryAddr == "" {
		pollFactoryAddr = os.Getenv("POLL_FACTORY_ADDRESS")
		if pollFactoryAddr == "" {
			return fmt.Errorf("factory address not specified (use --factory flag or set POLL_FACTORY_ADDRESS env var)")
		}
	}

	// If no voter merkle root specified, use a test root
	// NOTE: The contract requires a non-zero merkle root. For testing, we use keccak256("test-voters")
	if voterMerkleRoot == "" {
		voterMerkleRoot = "0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658" // keccak256("test")
		log.Println("No voter merkle root specified, using test root (for development only)")
	}

	log.Printf("  Question: %s\n", question)
	log.Printf("  Options: %v\n", optionsList)
	log.Printf("  Duration: %d seconds\n", duration)
	log.Printf("  Voter Merkle Root: %s\n", voterMerkleRoot)
	log.Printf("  PollFactory: %s\n", pollFactoryAddr)

	// Create PollFactory instance
	factoryAddr := common.HexToAddress(pollFactoryAddr)
	factory, err := bindings.NewPollFactory(factoryAddr, w.GetClient())
	if err != nil {
		return fmt.Errorf("failed to create factory instance: %w", err)
	}

	// Get auth
	auth, err := w.GetAuth(ctx)
	if err != nil {
		return fmt.Errorf("failed to get auth: %w", err)
	}

	log.Printf("  Gas Limit: %d\n", auth.GasLimit)
	log.Printf("  Gas Price: %s wei\n", auth.GasPrice.String())

	// Parse voter merkle root
	var merkleRoot [32]byte
	rootBytes := common.FromHex(voterMerkleRoot)
	if len(rootBytes) != 32 {
		return fmt.Errorf("invalid merkle root length: expected 32 bytes, got %d", len(rootBytes))
	}
	copy(merkleRoot[:], rootBytes)

	// Call createPoll on the contract
	log.Println("\nCalling PollFactory.createPoll()...")
	tx, err := factory.CreatePoll(auth, question, optionsList, big.NewInt(duration), merkleRoot)
	if err != nil {
		return fmt.Errorf("failed to create poll: %w", err)
	}

	log.Printf("Transaction submitted: %s\n", tx.Hash().Hex())
	log.Println("Waiting for confirmation...")

	// Wait for transaction receipt
	receipt, err := w.WaitForReceipt(ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("failed to get receipt: %w", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	// Parse logs to get the poll address from PollCreated event
	var pollAddress common.Address
	for _, vLog := range receipt.Logs {
		// Try to parse as PollCreated event
		event, err := factory.ParsePollCreated(*vLog)
		if err == nil {
			pollAddress = event.PollAddress
			break
		}
	}

	log.Printf("\nPoll created successfully!\n")
	log.Printf("Poll Address: %s\n", pollAddress.Hex())
	log.Printf("Transaction: %s\n", tx.Hash().Hex())
	log.Printf("Block Number: %d\n", receipt.BlockNumber.Uint64())
	log.Printf("Gas Used: %d\n", receipt.GasUsed)
	log.Printf("Closes at: %s\n", time.Now().Add(time.Duration(duration)*time.Second).Format(time.RFC3339))

	return nil
}
