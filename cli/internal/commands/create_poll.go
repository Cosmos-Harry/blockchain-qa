package commands

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/wallet"
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
	createPollCmd.Flags().StringVar(&voterMerkleRoot, "voter-root", "", "Merkle root of eligible voters")
	createPollCmd.Flags().StringVar(&pollFactoryAddr, "factory", "", "PollFactory contract address")

	createPollCmd.MarkFlagRequired("question")
	createPollCmd.MarkFlagRequired("options")
	createPollCmd.MarkFlagRequired("voter-root")
	createPollCmd.MarkFlagRequired("factory")
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

	// TODO: Call PollFactory.createPoll() when contract bindings are generated
	// For now, simulate the call
	log.Println("Simulating createPoll transaction...")
	log.Printf("  Question: %s\n", question)
	log.Printf("  Options: %v\n", optionsList)
	log.Printf("  Duration: %d seconds\n", duration)
	log.Printf("  Voter Merkle Root: %s\n", voterMerkleRoot)
	log.Printf("  PollFactory: %s\n", pollFactoryAddr)

	// Get auth
	auth, err := w.GetAuth(ctx)
	if err != nil {
		return fmt.Errorf("failed to get auth: %w", err)
	}

	log.Printf("  Gas Limit: %d\n", auth.GasLimit)
	log.Printf("  Gas Price: %s wei\n", auth.GasPrice.String())

	// Simulate successful transaction
	pollAddress := "0x1234567890123456789012345678901234567890"
	log.Printf("\nPoll created successfully!\n")
	log.Printf("Poll Address: %s\n", pollAddress)
	log.Printf("Closes at: %s\n", time.Now().Add(time.Duration(duration)*time.Second).Format(time.RFC3339))

	return nil
}
