package wallet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Wallet manages Ethereum account interactions
type Wallet struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    common.Address
	chainID    *big.Int
}

// NewReadOnlyWallet creates a wallet for read-only operations (no private key needed)
func NewReadOnlyWallet(ctx context.Context) (*Wallet, error) {
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

	return &Wallet{
		client:  client,
		chainID: chainID,
	}, nil
}

// NewWallet creates a new wallet from a private key
func NewWallet(ctx context.Context, privateKeyHex string) (*Wallet, error) {
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

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Wallet{
		client:     client,
		privateKey: privateKey,
		publicKey:  publicKeyECDSA,
		address:    address,
		chainID:    chainID,
	}, nil
}

// Address returns the wallet's Ethereum address
func (w *Wallet) Address() common.Address {
	return w.address
}

// GetAuth returns a transaction auth
func (w *Wallet) GetAuth(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(w.privateKey, w.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	nonce, err := w.client.PendingNonceAt(ctx, w.address)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	gasPrice, err := w.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)

	return auth, nil
}

// GetClient returns the Ethereum client
func (w *Wallet) GetClient() *ethclient.Client {
	return w.client
}

// Close closes the wallet's client connection
func (w *Wallet) Close() {
	w.client.Close()
}

// GetBalance returns the wallet's ETH balance
func (w *Wallet) GetBalance(ctx context.Context) (*big.Int, error) {
	return w.client.BalanceAt(ctx, w.address, nil)
}

// WaitForReceipt waits for a transaction receipt
func (w *Wallet) WaitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	// Wait for transaction to be mined
	for i := 0; i < 60; i++ { // Wait up to 2 minutes (60 * 2s)
		receipt, err := w.client.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		// If error is not "not found", return it
		if err.Error() != "not found" {
			time.Sleep(2 * time.Second)
			continue
		}

		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("transaction not mined after 2 minutes")
}
