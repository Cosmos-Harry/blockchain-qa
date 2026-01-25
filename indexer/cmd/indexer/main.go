package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cosmos-Harry/blockchain-qa/indexer/internal/blockchain"
	"github.com/Cosmos-Harry/blockchain-qa/indexer/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to database
	db, err := database.NewDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	// Connect to Ethereum node
	client, err := blockchain.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	log.Printf("Connected to Ethereum node (Chain ID: %s)\n", client.ChainID.String())

	// Get PollFactory address from environment
	pollFactory := os.Getenv("POLL_FACTORY_ADDRESS")
	if pollFactory == "" {
		log.Fatal("POLL_FACTORY_ADDRESS environment variable is required")
	}

	// Get start block from environment (default to 0)
	startBlock := uint64(0)
	// TODO: Parse START_BLOCK from environment if provided

	// Create and start event listener
	listener := blockchain.NewListener(client, db, pollFactory, startBlock)

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start listener in goroutine
	errChan := make(chan error, 1)
	go func() {
		errChan <- listener.Start(ctx)
	}()

	// Wait for shutdown signal or error
	select {
	case <-quit:
		log.Println("Received shutdown signal")
		cancel()
	case err := <-errChan:
		if err != nil && err != context.Canceled {
			log.Printf("Listener error: %v\n", err)
		}
	}

	log.Println("Indexer stopped")
}
