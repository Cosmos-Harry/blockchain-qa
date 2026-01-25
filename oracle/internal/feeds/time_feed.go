package feeds

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/Cosmos-Harry/blockchain-qa/oracle/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TimeFeed monitors polls and triggers closures at deadlines
type TimeFeed struct {
	client       *ethclient.Client
	auth         *bind.TransactOpts
	mode         types.ResponseMode
	pollRequests map[string]*types.PollCloseRequest
}

// NewTimeFeed creates a new time-based oracle feed
func NewTimeFeed(client *ethclient.Client, auth *bind.TransactOpts, mode types.ResponseMode) *TimeFeed {
	return &TimeFeed{
		client:       client,
		auth:         auth,
		mode:         mode,
		pollRequests: make(map[string]*types.PollCloseRequest),
	}
}

// SetMode changes the response mode
func (f *TimeFeed) SetMode(mode types.ResponseMode) {
	f.mode = mode
	log.Printf("Oracle mode set to: %s\n", mode.String())
}

// RegisterPollClose registers a poll for closing at its deadline
func (f *TimeFeed) RegisterPollClose(pollAddress string, deadline time.Time) {
	f.pollRequests[pollAddress] = &types.PollCloseRequest{
		PollAddress: pollAddress,
		Deadline:    deadline,
		RequestedAt: time.Now(),
	}
	log.Printf("Registered poll %s for closing at %s\n", pollAddress, deadline.Format(time.RFC3339))
}

// Start begins monitoring and closing polls
func (f *TimeFeed) Start(ctx context.Context) error {
	log.Printf("Starting oracle time feed in %s mode\n", f.mode.String())

	ticker := time.NewTicker(10 * time.Second) // Check every 10 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping oracle time feed...")
			return ctx.Err()
		case <-ticker.C:
			f.checkAndClosePolls(ctx)
		}
	}
}

// checkAndClosePolls checks for polls that should be closed
func (f *TimeFeed) checkAndClosePolls(ctx context.Context) {
	now := time.Now()

	for pollAddress, request := range f.pollRequests {
		shouldClose, delay := f.shouldClosePoll(request, now)
		if !shouldClose {
			continue
		}

		// Apply delay based on mode
		if delay > 0 {
			log.Printf("Delaying poll closure by %v (mode: %s)\n", delay, f.mode.String())
			time.Sleep(delay)
		}

		if err := f.closePoll(ctx, pollAddress); err != nil {
			log.Printf("Error closing poll %s: %v\n", pollAddress, err)
			continue
		}

		// Remove from pending requests
		delete(f.pollRequests, pollAddress)
		log.Printf("Successfully closed poll %s\n", pollAddress)
	}
}

// shouldClosePoll determines if a poll should be closed based on mode
func (f *TimeFeed) shouldClosePoll(request *types.PollCloseRequest, now time.Time) (bool, time.Duration) {
	switch f.mode {
	case types.OnTime:
		// Close exactly at deadline
		if now.After(request.Deadline) {
			return true, 0
		}
		return false, 0

	case types.Late:
		// Close 5-15 minutes after deadline
		lateDelay := time.Duration(5+rand.Intn(10)) * time.Minute
		closeTime := request.Deadline.Add(lateDelay)
		if now.After(closeTime) {
			return true, 0
		}
		return false, 0

	case types.Invalid:
		// Still close on time, but with invalid data
		if now.After(request.Deadline) {
			return true, 0
		}
		return false, 0

	case types.NoResponse:
		// Never close
		return false, 0

	default:
		return false, 0
	}
}

// closePoll sends a transaction to close a poll
func (f *TimeFeed) closePoll(ctx context.Context, pollAddress string) error {
	// In a real implementation, this would call the oracle contract's
	// requestPollClose or fulfillRequest function

	// For now, we simulate the transaction
	log.Printf("Sending closePoll transaction for poll %s\n", pollAddress)

	// Simulate transaction
	pollAddr := common.HexToAddress(pollAddress)

	// Get current nonce
	nonce, err := f.client.PendingNonceAt(ctx, f.auth.From)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get gas price
	gasPrice, err := f.client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to get gas price: %w", err)
	}

	// In Invalid mode, send corrupted data
	if f.mode == types.Invalid {
		log.Println("Warning: Sending invalid data (mode: Invalid)")
		gasPrice = big.NewInt(0) // Invalid gas price
	}

	// Update auth
	f.auth.Nonce = big.NewInt(int64(nonce))
	f.auth.GasPrice = gasPrice
	f.auth.GasLimit = uint64(300000)

	// TODO: Call actual contract method when contract bindings are generated
	// For now, just log the action
	log.Printf("Would call oracle.requestPollClose(%s) with nonce=%d, gasPrice=%s\n",
		pollAddr.Hex(), nonce, gasPrice.String())

	return nil
}
