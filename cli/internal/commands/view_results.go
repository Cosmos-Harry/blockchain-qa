package commands

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/bindings"
	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var viewResultsCmd = &cobra.Command{
	Use:   "view-results",
	Short: "View poll results",
	Long:  `Query the poll contract directly to view poll details and results.`,
	RunE:  runViewResults,
}

func init() {
	rootCmd.AddCommand(viewResultsCmd)

	viewResultsCmd.Flags().StringVar(&pollAddress, "poll", "", "Poll contract address")
	viewResultsCmd.MarkFlagRequired("poll")
}

func runViewResults(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	log.Printf("Fetching results for poll: %s\n", pollAddress)

	// Create read-only wallet (no private key needed for view functions)
	w, err := wallet.NewReadOnlyWallet(ctx)
	if err != nil {
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	defer w.Close()

	// Create Poll contract instance
	pollAddr := common.HexToAddress(pollAddress)
	poll, err := bindings.NewPoll(pollAddr, w.GetClient())
	if err != nil {
		return fmt.Errorf("failed to create poll instance: %w", err)
	}

	// Fetch poll details from contract
	question, err := poll.Question(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch question: %w", err)
	}

	options, err := poll.PollOptions(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch options: %w", err)
	}

	stateUint, err := poll.State(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch state: %w", err)
	}

	createdAt, err := poll.CreatedAt(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch created time: %w", err)
	}

	endTime, err := poll.EndTime(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch end time: %w", err)
	}

	totalCommitted, err := poll.TotalCommitted(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch total committed: %w", err)
	}

	totalRevealed, err := poll.TotalRevealed(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch total revealed: %w", err)
	}

	// Convert state to string
	var stateStr string
	switch stateUint {
	case 0:
		stateStr = "Active"
	case 1:
		stateStr = "Closed"
	case 2:
		stateStr = "Tallied"
	default:
		stateStr = "Unknown"
	}

	// Display poll details
	fmt.Println("\n=== Poll Details ===")
	fmt.Printf("Question: %s\n", question)
	fmt.Printf("State: %s\n", stateStr)
	fmt.Printf("Created: %s\n", time.Unix(createdAt.Int64(), 0).Format(time.RFC3339))
	fmt.Printf("Closes: %s\n", time.Unix(endTime.Int64(), 0).Format(time.RFC3339))

	fmt.Println("\n=== Options ===")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i, option)
	}

	// Display vote statistics
	fmt.Println("\n=== Vote Statistics ===")
	fmt.Printf("Total Committed: %d\n", totalCommitted.Uint64())
	fmt.Printf("Total Revealed: %d\n", totalRevealed.Uint64())
	fmt.Printf("Pending Reveals: %d\n", totalCommitted.Uint64()-totalRevealed.Uint64())

	// If tallied, fetch and display results
	if stateUint == 2 { // Tallied
		results, err := poll.GetResults(nil)
		if err != nil {
			log.Printf("Warning: failed to fetch results: %v\n", err)
		} else {
			fmt.Println("\n=== Results ===")
			for i, count := range results {
				if i < len(options) {
					fmt.Printf("%s: %d votes\n", options[i], count.Uint64())
				}
			}
		}
	} else {
		fmt.Println("\nResults not yet tallied (poll must be closed and votes revealed)")
	}

	return nil
}
