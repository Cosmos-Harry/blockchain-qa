package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// CreatePoll inserts a new poll into the database
func (db *DB) CreatePoll(ctx context.Context, poll *Poll) error {
	query := `
		INSERT INTO polls (
			contract_address, question, options, duration, voter_merkle_root,
			created_at, closes_at, state, creator, block_number, transaction_hash
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, created_timestamp
	`

	err := db.Pool.QueryRow(
		ctx, query,
		poll.ContractAddress, poll.Question, poll.Options, poll.Duration,
		poll.VoterMerkleRoot, poll.CreatedAt, poll.ClosesAt, poll.State,
		poll.Creator, poll.BlockNumber, poll.TransactionHash,
	).Scan(&poll.ID, &poll.CreatedTimestamp)

	if err != nil {
		return fmt.Errorf("failed to create poll: %w", err)
	}

	return nil
}

// GetPollByAddress retrieves a poll by its contract address
func (db *DB) GetPollByAddress(ctx context.Context, address string) (*Poll, error) {
	query := `
		SELECT id, contract_address, question, options, duration, voter_merkle_root,
			created_at, closes_at, state, creator, block_number, transaction_hash,
			created_timestamp
		FROM polls
		WHERE contract_address = $1
	`

	poll := &Poll{}
	err := db.Pool.QueryRow(ctx, query, address).Scan(
		&poll.ID, &poll.ContractAddress, &poll.Question, &poll.Options,
		&poll.Duration, &poll.VoterMerkleRoot, &poll.CreatedAt, &poll.ClosesAt,
		&poll.State, &poll.Creator, &poll.BlockNumber, &poll.TransactionHash,
		&poll.CreatedTimestamp,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get poll: %w", err)
	}

	return poll, nil
}

// ListPolls retrieves all polls with optional state filter
func (db *DB) ListPolls(ctx context.Context, state string, limit, offset int) ([]*Poll, error) {
	var query string
	var args []interface{}

	if state != "" {
		query = `
			SELECT id, contract_address, question, options, duration, voter_merkle_root,
				created_at, closes_at, state, creator, block_number, transaction_hash,
				created_timestamp
			FROM polls
			WHERE state = $1
			ORDER BY created_at DESC
			LIMIT $2 OFFSET $3
		`
		args = []interface{}{state, limit, offset}
	} else {
		query = `
			SELECT id, contract_address, question, options, duration, voter_merkle_root,
				created_at, closes_at, state, creator, block_number, transaction_hash,
				created_timestamp
			FROM polls
			ORDER BY created_at DESC
			LIMIT $1 OFFSET $2
		`
		args = []interface{}{limit, offset}
	}

	rows, err := db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list polls: %w", err)
	}
	defer rows.Close()

	var polls []*Poll
	for rows.Next() {
		poll := &Poll{}
		err := rows.Scan(
			&poll.ID, &poll.ContractAddress, &poll.Question, &poll.Options,
			&poll.Duration, &poll.VoterMerkleRoot, &poll.CreatedAt, &poll.ClosesAt,
			&poll.State, &poll.Creator, &poll.BlockNumber, &poll.TransactionHash,
			&poll.CreatedTimestamp,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan poll: %w", err)
		}
		polls = append(polls, poll)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return polls, nil
}

// UpdatePollState updates the state of a poll
func (db *DB) UpdatePollState(ctx context.Context, address, state string) error {
	query := `UPDATE polls SET state = $1 WHERE contract_address = $2`

	result, err := db.Pool.Exec(ctx, query, state, address)
	if err != nil {
		return fmt.Errorf("failed to update poll state: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("poll not found: %s", address)
	}

	return nil
}
