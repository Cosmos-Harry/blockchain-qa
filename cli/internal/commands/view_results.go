package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	apiURL string
)

var viewResultsCmd = &cobra.Command{
	Use:   "view-results",
	Short: "View poll results",
	Long:  `Query the indexer API to view poll results and statistics.`,
	RunE:  runViewResults,
}

func init() {
	rootCmd.AddCommand(viewResultsCmd)

	viewResultsCmd.Flags().StringVar(&pollAddress, "poll", "", "Poll contract address")
	viewResultsCmd.Flags().StringVar(&apiURL, "api-url", "http://localhost:3000", "Indexer API URL")

	viewResultsCmd.MarkFlagRequired("poll")
}

func runViewResults(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	log.Printf("Fetching results for poll: %s\n", pollAddress)

	// Query poll details
	poll, err := fetchPollDetails(ctx, pollAddress)
	if err != nil {
		return fmt.Errorf("failed to fetch poll details: %w", err)
	}

	fmt.Println("\n=== Poll Details ===")
	fmt.Printf("Question: %s\n", poll.Question)
	fmt.Printf("State: %s\n", poll.State)
	fmt.Printf("Created: %s\n", poll.CreatedAt)
	fmt.Printf("Closes: %s\n", poll.ClosesAt)

	// Query vote statistics
	stats, err := fetchVoteStats(ctx, pollAddress)
	if err != nil {
		return fmt.Errorf("failed to fetch vote stats: %w", err)
	}

	fmt.Println("\n=== Vote Statistics ===")
	fmt.Printf("Total Votes: %d\n", stats.TotalVotes)
	fmt.Printf("Revealed Votes: %d\n", stats.RevealedVotes)
	fmt.Printf("Pending Reveals: %d\n", stats.PendingReveals)

	// Query results if tallied
	if poll.State == "tallied" {
		results, err := fetchResults(ctx, pollAddress)
		if err != nil {
			log.Printf("Warning: failed to fetch results: %v\n", err)
		} else {
			fmt.Println("\n=== Results ===")
			for i, count := range results.VoteCounts {
				if i < len(poll.Options) {
					fmt.Printf("%s: %d votes\n", poll.Options[i], count)
				}
			}
		}
	} else {
		fmt.Println("\nResults not yet tallied (poll must be closed and votes revealed)")
	}

	return nil
}

type PollDetails struct {
	Question  string   `json:"question"`
	Options   []string `json:"options"`
	State     string   `json:"state"`
	CreatedAt string   `json:"created_at"`
	ClosesAt  string   `json:"closes_at"`
}

type VoteStats struct {
	TotalVotes     int `json:"total_votes"`
	RevealedVotes  int `json:"revealed_votes"`
	PendingReveals int `json:"pending_reveals"`
}

type Results struct {
	VoteCounts []int `json:"vote_counts"`
	TotalVotes int   `json:"total_votes"`
}

func fetchPollDetails(ctx context.Context, pollAddr string) (*PollDetails, error) {
	url := fmt.Sprintf("%s/api/polls/%s", apiURL, pollAddr)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var poll PollDetails
	if err := json.NewDecoder(resp.Body).Decode(&poll); err != nil {
		return nil, err
	}

	return &poll, nil
}

func fetchVoteStats(ctx context.Context, pollAddr string) (*VoteStats, error) {
	url := fmt.Sprintf("%s/api/polls/%s/stats", apiURL, pollAddr)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var stats VoteStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

func fetchResults(ctx context.Context, pollAddr string) (*Results, error) {
	url := fmt.Sprintf("%s/api/polls/%s/results", apiURL, pollAddr)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("results not available")
	}

	var results Results
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return &results, nil
}

func init() {
	// Override environment variables with flags
	if rpcURL != "" {
		os.Setenv("RPC_URL", rpcURL)
	}
}
