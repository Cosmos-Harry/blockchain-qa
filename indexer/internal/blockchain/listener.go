package blockchain

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/Cosmos-Harry/blockchain-qa/indexer/internal/database"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Listener listens for blockchain events and processes them
type Listener struct {
	client   *Client
	db       *database.DB
	pollFactory common.Address
	startBlock uint64
}

// NewListener creates a new event listener
func NewListener(client *Client, db *database.DB, pollFactory string, startBlock uint64) *Listener {
	return &Listener{
		client:      client,
		db:          db,
		pollFactory: common.HexToAddress(pollFactory),
		startBlock:  startBlock,
	}
}

// Start begins listening for events
func (l *Listener) Start(ctx context.Context) error {
	log.Println("Starting blockchain event listener...")

	// Get last processed block from database
	lastBlock, err := l.db.GetLastProcessedBlock(ctx)
	if err != nil {
		return fmt.Errorf("failed to get last processed block: %w", err)
	}

	if lastBlock > 0 {
		l.startBlock = uint64(lastBlock) + 1
		log.Printf("Resuming from block %d\n", l.startBlock)
	} else {
		log.Printf("Starting from block %d\n", l.startBlock)
	}

	// Subscribe to new blocks
	headers := make(chan *types.Header)
	sub, err := l.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new blocks: %w", err)
	}
	defer sub.Unsubscribe()

	// Process historical blocks first
	if err := l.processHistoricalBlocks(ctx); err != nil {
		log.Printf("Warning: failed to process historical blocks: %v\n", err)
	}

	// Process new blocks as they arrive
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping event listener...")
			return ctx.Err()
		case err := <-sub.Err():
			return fmt.Errorf("subscription error: %w", err)
		case header := <-headers:
			if err := l.processBlock(ctx, header.Number.Uint64()); err != nil {
				log.Printf("Error processing block %d: %v\n", header.Number.Uint64(), err)
			}
		}
	}
}

// processHistoricalBlocks processes all blocks from startBlock to current block
func (l *Listener) processHistoricalBlocks(ctx context.Context) error {
	currentBlock, err := l.client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get current block number: %w", err)
	}

	if l.startBlock >= currentBlock {
		return nil // No historical blocks to process
	}

	log.Printf("Processing historical blocks from %d to %d\n", l.startBlock, currentBlock)

	// Process in batches to avoid overwhelming the node
	batchSize := uint64(1000)
	for fromBlock := l.startBlock; fromBlock < currentBlock; fromBlock += batchSize {
		toBlock := fromBlock + batchSize - 1
		if toBlock > currentBlock {
			toBlock = currentBlock
		}

		if err := l.processBlockRange(ctx, fromBlock, toBlock); err != nil {
			return fmt.Errorf("failed to process block range %d-%d: %w", fromBlock, toBlock, err)
		}

		log.Printf("Processed blocks %d-%d\n", fromBlock, toBlock)
	}

	return nil
}

// processBlockRange processes a range of blocks
func (l *Listener) processBlockRange(ctx context.Context, fromBlock, toBlock uint64) error {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []common.Address{l.pollFactory},
	}

	logs, err := l.client.FilterLogs(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to filter logs: %w", err)
	}

	for _, vLog := range logs {
		if err := l.processLog(ctx, vLog); err != nil {
			log.Printf("Error processing log: %v\n", err)
		}
	}

	return nil
}

// processBlock processes a single block
func (l *Listener) processBlock(ctx context.Context, blockNumber uint64) error {
	return l.processBlockRange(ctx, blockNumber, blockNumber)
}

// processLog processes a single log entry
func (l *Listener) processLog(ctx context.Context, vLog types.Log) error {
	// Store raw event
	eventData, err := json.Marshal(map[string]interface{}{
		"topics": vLog.Topics,
		"data":   common.Bytes2Hex(vLog.Data),
	})
	if err != nil {
		return fmt.Errorf("failed to marshal event data: %w", err)
	}

	event := &database.Event{
		ContractAddress: vLog.Address.Hex(),
		EventName:       "Unknown", // Will be determined by topic
		EventData:       string(eventData),
		BlockNumber:     int64(vLog.BlockNumber),
		BlockHash:       vLog.BlockHash.Hex(),
		TransactionHash: vLog.TxHash.Hex(),
		LogIndex:        int(vLog.Index),
	}

	// Determine event type by topic
	if len(vLog.Topics) > 0 {
		event.EventName = l.getEventNameByTopic(vLog.Topics[0])
	}

	if err := l.db.CreateEvent(ctx, event); err != nil {
		return fmt.Errorf("failed to save event: %w", err)
	}

	// Process specific event types
	switch event.EventName {
	case "PollCreated":
		return l.processPollCreatedEvent(ctx, vLog)
	case "VoteCommitted":
		return l.processVoteCommittedEvent(ctx, vLog)
	case "VoteRevealed":
		return l.processVoteRevealedEvent(ctx, vLog)
	case "PollClosed":
		return l.processPollClosedEvent(ctx, vLog)
	case "ResultsTallied":
		return l.processResultsTalliedEvent(ctx, vLog)
	}

	return nil
}

// getEventNameByTopic returns the event name for a given topic hash
func (l *Listener) getEventNameByTopic(topic common.Hash) string {
	// Event signatures (keccak256 hash of event signature)
	// These will be replaced with actual ABI-generated signatures
	switch topic.Hex() {
	case "0x5c...": // PollCreated signature (placeholder)
		return "PollCreated"
	case "0x8a...": // VoteCommitted signature (placeholder)
		return "VoteCommitted"
	case "0x7d...": // VoteRevealed signature (placeholder)
		return "VoteRevealed"
	case "0x6b...": // PollClosed signature (placeholder)
		return "PollClosed"
	case "0x4e...": // ResultsTallied signature (placeholder)
		return "ResultsTallied"
	default:
		return "Unknown"
	}
}

// processPollCreatedEvent processes a PollCreated event
func (l *Listener) processPollCreatedEvent(ctx context.Context, vLog types.Log) error {
	// TODO: Parse event data using ABI
	// For now, this is a placeholder
	log.Printf("Processing PollCreated event at block %d\n", vLog.BlockNumber)
	return nil
}

// processVoteCommittedEvent processes a VoteCommitted event
func (l *Listener) processVoteCommittedEvent(ctx context.Context, vLog types.Log) error {
	// TODO: Parse event data using ABI
	log.Printf("Processing VoteCommitted event at block %d\n", vLog.BlockNumber)
	return nil
}

// processVoteRevealedEvent processes a VoteRevealed event
func (l *Listener) processVoteRevealedEvent(ctx context.Context, vLog types.Log) error {
	// TODO: Parse event data using ABI
	log.Printf("Processing VoteRevealed event at block %d\n", vLog.BlockNumber)
	return nil
}

// processPollClosedEvent processes a PollClosed event
func (l *Listener) processPollClosedEvent(ctx context.Context, vLog types.Log) error {
	// TODO: Parse event data using ABI
	log.Printf("Processing PollClosed event at block %d\n", vLog.BlockNumber)

	// Update poll state to closed
	pollAddress := vLog.Address.Hex()
	return l.db.UpdatePollState(ctx, pollAddress, "closed")
}

// processResultsTalliedEvent processes a ResultsTallied event
func (l *Listener) processResultsTalliedEvent(ctx context.Context, vLog types.Log) error {
	// TODO: Parse event data using ABI
	log.Printf("Processing ResultsTallied event at block %d\n", vLog.BlockNumber)

	// Update poll state to tallied
	pollAddress := vLog.Address.Hex()
	return l.db.UpdatePollState(ctx, pollAddress, "tallied")
}
