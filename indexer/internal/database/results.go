package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// CreateResult inserts tallied poll results into the database
func (db *DB) CreateResult(ctx context.Context, result *Result) error {
	query := `
		INSERT INTO results (
			poll_address, vote_counts, total_votes, tallied_at, block_number, transaction_hash
		) VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (poll_address) DO UPDATE
		SET vote_counts = EXCLUDED.vote_counts,
			total_votes = EXCLUDED.total_votes,
			tallied_at = EXCLUDED.tallied_at,
			block_number = EXCLUDED.block_number,
			transaction_hash = EXCLUDED.transaction_hash
		RETURNING id, created_timestamp
	`

	err := db.Pool.QueryRow(
		ctx, query,
		result.PollAddress, result.VoteCounts, result.TotalVotes,
		result.TalliedAt, result.BlockNumber, result.TransactionHash,
	).Scan(&result.ID, &result.CreatedTimestamp)

	if err != nil {
		return fmt.Errorf("failed to create result: %w", err)
	}

	return nil
}

// GetResultByPoll retrieves the results for a specific poll
func (db *DB) GetResultByPoll(ctx context.Context, pollAddress string) (*Result, error) {
	query := `
		SELECT id, poll_address, vote_counts, total_votes, tallied_at,
			block_number, transaction_hash, created_timestamp
		FROM results
		WHERE poll_address = $1
	`

	result := &Result{}
	err := db.Pool.QueryRow(ctx, query, pollAddress).Scan(
		&result.ID, &result.PollAddress, &result.VoteCounts, &result.TotalVotes,
		&result.TalliedAt, &result.BlockNumber, &result.TransactionHash,
		&result.CreatedTimestamp,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get result: %w", err)
	}

	return result, nil
}
