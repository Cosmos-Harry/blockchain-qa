package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Cosmos-Harry/blockchain-qa/oracle/internal/feeds"
	"github.com/Cosmos-Harry/blockchain-qa/oracle/internal/publisher"
	"github.com/Cosmos-Harry/blockchain-qa/oracle/internal/types"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create transaction publisher
	pub, err := publisher.NewPublisher(ctx)
	if err != nil {
		log.Fatalf("Failed to create publisher: %v", err)
	}
	defer pub.Close()

	// Get oracle mode from environment
	mode := getOracleModeFromEnv()
	log.Printf("Starting oracle in %s mode\n", mode.String())

	// Create time feed
	timeFeed := feeds.NewTimeFeed(pub.GetClient(), pub.GetAuth(), mode)

	// Register some test polls for demonstration
	// In production, this would query the indexer API or listen to PollCreated events
	registerTestPolls(timeFeed)

	// Start time feed in goroutine
	errChan := make(chan error, 1)
	go func() {
		errChan <- timeFeed.Start(ctx)
	}()

	// Handle mode changes via signals (for testing)
	// In production, this might be an HTTP API endpoint
	go handleModeChanges(timeFeed)

	// Wait for shutdown signal or error
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("Received shutdown signal")
		cancel()
	case err := <-errChan:
		if err != nil && err != context.Canceled {
			log.Printf("Oracle error: %v\n", err)
		}
	}

	log.Println("Oracle stopped")
}

// getOracleModeFromEnv returns the oracle mode from environment variable
func getOracleModeFromEnv() types.ResponseMode {
	modeStr := os.Getenv("ORACLE_MODE")
	switch modeStr {
	case "late", "Late", "LATE":
		return types.Late
	case "invalid", "Invalid", "INVALID":
		return types.Invalid
	case "noresponse", "NoResponse", "NO_RESPONSE":
		return types.NoResponse
	default:
		return types.OnTime
	}
}

// registerTestPolls registers some test polls for demonstration
func registerTestPolls(feed *feeds.TimeFeed) {
	// Read test poll addresses from environment
	pollAddresses := os.Getenv("TEST_POLL_ADDRESSES")
	if pollAddresses == "" {
		log.Println("No test polls configured (set TEST_POLL_ADDRESSES)")
		return
	}

	// Read deadline offset in seconds
	deadlineOffsetStr := os.Getenv("TEST_POLL_DEADLINE_OFFSET")
	deadlineOffset := 300 // Default 5 minutes
	if deadlineOffsetStr != "" {
		if offset, err := strconv.Atoi(deadlineOffsetStr); err == nil {
			deadlineOffset = offset
		}
	}

	deadline := time.Now().Add(time.Duration(deadlineOffset) * time.Second)

	// Register each poll
	// In production, this would be a comma-separated list
	feed.RegisterPollClose(pollAddresses, deadline)
	log.Printf("Registered test poll %s with deadline in %d seconds\n", pollAddresses, deadlineOffset)
}

// handleModeChanges allows changing oracle mode at runtime (for testing)
func handleModeChanges(feed *feeds.TimeFeed) {
	// This is a simple implementation for testing
	// In production, this might be an HTTP endpoint or configuration reload
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Check if mode change is requested via environment
		newModeStr := os.Getenv("ORACLE_MODE_CHANGE")
		if newModeStr == "" {
			continue
		}

		var newMode types.ResponseMode
		switch newModeStr {
		case "ontime", "OnTime", "ON_TIME":
			newMode = types.OnTime
		case "late", "Late", "LATE":
			newMode = types.Late
		case "invalid", "Invalid", "INVALID":
			newMode = types.Invalid
		case "noresponse", "NoResponse", "NO_RESPONSE":
			newMode = types.NoResponse
		default:
			continue
		}

		feed.SetMode(newMode)
		// Clear the change request
		os.Unsetenv("ORACLE_MODE_CHANGE")
	}
}
