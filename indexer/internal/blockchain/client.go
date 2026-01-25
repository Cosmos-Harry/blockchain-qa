package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Client wraps an Ethereum client connection
type Client struct {
	*ethclient.Client
	ChainID *big.Int
}

// NewClient creates a new Ethereum client
func NewClient(ctx context.Context) (*Client, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "http://localhost:8545" // Default to local Anvil
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	return &Client{
		Client:  client,
		ChainID: chainID,
	}, nil
}
