package publisher

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Publisher handles transaction publishing to the blockchain
type Publisher struct {
	client  *ethclient.Client
	auth    *bind.TransactOpts
	chainID *big.Int
}

// NewPublisher creates a new transaction publisher
func NewPublisher(ctx context.Context) (*Publisher, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "http://localhost:8545"
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Get private key from environment
	privateKeyHex := os.Getenv("ORACLE_PRIVATE_KEY")
	if privateKeyHex == "" {
		return nil, fmt.Errorf("ORACLE_PRIVATE_KEY environment variable is required")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Create transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("Oracle publisher address: %s\n", fromAddress.Hex())

	return &Publisher{
		client:  client,
		auth:    auth,
		chainID: chainID,
	}, nil
}

// GetAuth returns the transaction auth
func (p *Publisher) GetAuth() *bind.TransactOpts {
	return p.auth
}

// GetClient returns the Ethereum client
func (p *Publisher) GetClient() *ethclient.Client {
	return p.client
}

// Close closes the Ethereum client connection
func (p *Publisher) Close() {
	p.client.Close()
}
