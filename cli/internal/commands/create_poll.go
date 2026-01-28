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

	// Get factory address from flag, environment, or saved state
	state := loadState()
	if pollFactoryAddr == "" {
		pollFactoryAddr = os.Getenv("POLL_FACTORY_ADDRESS")
	}
	if pollFactoryAddr == "" {
		pollFactoryAddr = state.FactoryAddress
	}
	if pollFactoryAddr == "" {
		return fmt.Errorf("factory address not specified (use --factory flag, set POLL_FACTORY_ADDRESS env var, or save it to state)")
	}

	// If no voter merkle root specified, default to single-voter root for Anvil account 0.
	// This is keccak256(abi.encodePacked(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266)), allowing
	// that account to vote with an empty Merkle proof.
	if voterMerkleRoot == "" {
		voterMerkleRoot = "0xe9707d0e6171f728f7473c24cc0432a9b07eaaf1efed6a137a4a8c12c79552d9"
		log.Println("No voter merkle root specified, defaulting to Anvil account 0 single-voter root")
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

	// Save state for subsequent commands
	state.PollAddress = pollAddress.Hex()
	state.FactoryAddress = pollFactoryAddr
	state.Nonce = ""  // reset vote state for new poll
	state.Choice = 0
	if err := saveState(state); err != nil {
		log.Printf("Warning: failed to save state: %v\n", err)
	}

	log.Printf("\nPoll created successfully!\n")
	log.Printf("Poll Address: %s\n", pollAddress.Hex())
	log.Printf("Transaction: %s\n", tx.Hash().Hex())
	log.Printf("Block Number: %d\n", receipt.BlockNumber.Uint64())
	log.Printf("Gas Used: %d\n", receipt.GasUsed)
	log.Printf("Closes at: %s\n", time.Now().Add(time.Duration(duration)*time.Second).Format(time.RFC3339))
	log.Printf("State saved to %s\n", stateFileName)

	return nil
}
