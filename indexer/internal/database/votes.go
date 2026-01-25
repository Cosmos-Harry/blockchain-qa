package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// CreateVote inserts a new vote commitment into the database
func (db *DB) CreateVote(ctx context.Context, vote *Vote) error {
	query := `
		INSERT INTO votes (
			poll_address, voter, commitment, committed_at, block_number, transaction_hash
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_timestamp
	`

	err := db.Pool.QueryRow(
		ctx, query,
		vote.PollAddress, vote.Voter, vote.Commitment,
		vote.CommittedAt, vote.BlockNumber, vote.TransactionHash,
	).Scan(&vote.ID, &vote.CreatedTimestamp)

	if err != nil {
		return fmt.Errorf("failed to create vote: %w", err)
	}

	return nil
}

// RevealVote updates a vote with the revealed choice and nonce
func (db *DB) RevealVote(ctx context.Context, pollAddress, voter string, choice int, nonce []byte) error {
	query := `
		UPDATE votes
		SET choice = $3, nonce = $4, revealed = true, revealed_at = NOW()
		WHERE poll_address = $1 AND voter = $2
	`

	result, err := db.Pool.Exec(ctx, query, pollAddress, voter, choice, nonce)
	if err != nil {
		return fmt.Errorf("failed to reveal vote: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("vote not found for poll %s and voter %s", pollAddress, voter)
	}

	return nil
}

// GetVote retrieves a vote by poll address and voter
func (db *DB) GetVote(ctx context.Context, pollAddress, voter string) (*Vote, error) {
	query := `
		SELECT id, poll_address, voter, commitment, choice, nonce, revealed,
			committed_at, revealed_at, block_number, transaction_hash, created_timestamp
		FROM votes
		WHERE poll_address = $1 AND voter = $2
	`

	vote := &Vote{}
	err := db.Pool.QueryRow(ctx, query, pollAddress, voter).Scan(
		&vote.ID, &vote.PollAddress, &vote.Voter, &vote.Commitment,
		&vote.Choice, &vote.Nonce, &vote.Revealed, &vote.CommittedAt,
		&vote.RevealedAt, &vote.BlockNumber, &vote.TransactionHash,
		&vote.CreatedTimestamp,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get vote: %w", err)
	}

	return vote, nil
}

// ListVotesByPoll retrieves all votes for a specific poll
func (db *DB) ListVotesByPoll(ctx context.Context, pollAddress string, revealedOnly bool) ([]*Vote, error) {
	var query string
	if revealedOnly {
		query = `
			SELECT id, poll_address, voter, commitment, choice, nonce, revealed,
				committed_at, revealed_at, block_number, transaction_hash, created_timestamp
			FROM votes
			WHERE poll_address = $1 AND revealed = true
			ORDER BY committed_at ASC
		`
	} else {
		query = `
			SELECT id, poll_address, voter, commitment, choice, nonce, revealed,
				committed_at, revealed_at, block_number, transaction_hash, created_timestamp
			FROM votes
			WHERE poll_address = $1
			ORDER BY committed_at ASC
		`
	}

	rows, err := db.Pool.Query(ctx, query, pollAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to list votes: %w", err)
	}
	defer rows.Close()

	var votes []*Vote
	for rows.Next() {
		vote := &Vote{}
		err := rows.Scan(
			&vote.ID, &vote.PollAddress, &vote.Voter, &vote.Commitment,
			&vote.Choice, &vote.Nonce, &vote.Revealed, &vote.CommittedAt,
			&vote.RevealedAt, &vote.BlockNumber, &vote.TransactionHash,
			&vote.CreatedTimestamp,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vote: %w", err)
		}
		votes = append(votes, vote)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return votes, nil
}

// GetVoteCount returns the total number of votes for a poll
func (db *DB) GetVoteCount(ctx context.Context, pollAddress string, revealedOnly bool) (int, error) {
	var query string
	if revealedOnly {
		query = `SELECT COUNT(*) FROM votes WHERE poll_address = $1 AND revealed = true`
	} else {
		query = `SELECT COUNT(*) FROM votes WHERE poll_address = $1`
	}

	var count int
	err := db.Pool.QueryRow(ctx, query, pollAddress).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get vote count: %w", err)
	}

	return count, nil
}
