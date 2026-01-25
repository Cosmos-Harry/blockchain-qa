package database

import (
	"context"
	"fmt"
)

// CreateEvent inserts a new event into the database
func (db *DB) CreateEvent(ctx context.Context, event *Event) error {
	query := `
		INSERT INTO events (
			contract_address, event_name, event_data, block_number,
			block_hash, transaction_hash, log_index
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (transaction_hash, log_index) DO NOTHING
		RETURNING id, created_timestamp
	`

	err := db.Pool.QueryRow(
		ctx, query,
		event.ContractAddress, event.EventName, event.EventData,
		event.BlockNumber, event.BlockHash, event.TransactionHash,
		event.LogIndex,
	).Scan(&event.ID, &event.CreatedTimestamp)

	// Ignore duplicate key errors (event already processed)
	if err != nil && err.Error() != "no rows in result set" {
		return fmt.Errorf("failed to create event: %w", err)
	}

	return nil
}

// GetLastProcessedBlock retrieves the highest block number processed
func (db *DB) GetLastProcessedBlock(ctx context.Context) (int64, error) {
	query := `SELECT COALESCE(MAX(block_number), 0) FROM events`

	var blockNumber int64
	err := db.Pool.QueryRow(ctx, query).Scan(&blockNumber)
	if err != nil {
		return 0, fmt.Errorf("failed to get last processed block: %w", err)
	}

	return blockNumber, nil
}
