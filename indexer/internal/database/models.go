package database

import (
	"time"
)

// Poll represents a poll in the database
type Poll struct {
	ID                int       `json:"id"`
	ContractAddress   string    `json:"contract_address"`
	Question          string    `json:"question"`
	Options           []string  `json:"options"`
	Duration          int       `json:"duration"`
	VoterMerkleRoot   string    `json:"voter_merkle_root"`
	CreatedAt         time.Time `json:"created_at"`
	ClosesAt          time.Time `json:"closes_at"`
	State             string    `json:"state"`
	Creator           string    `json:"creator"`
	BlockNumber       int64     `json:"block_number"`
	TransactionHash   string    `json:"transaction_hash"`
	CreatedTimestamp  time.Time `json:"created_timestamp"`
}

// Vote represents a vote in the database
type Vote struct {
	ID              int        `json:"id"`
	PollAddress     string     `json:"poll_address"`
	Voter           string     `json:"voter"`
	Commitment      string     `json:"commitment"`
	Choice          *int       `json:"choice,omitempty"`
	Nonce           []byte     `json:"nonce,omitempty"`
	Revealed        bool       `json:"revealed"`
	CommittedAt     time.Time  `json:"committed_at"`
	RevealedAt      *time.Time `json:"revealed_at,omitempty"`
	BlockNumber     int64      `json:"block_number"`
	TransactionHash string     `json:"transaction_hash"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
}

// Event represents a blockchain event in the database
type Event struct {
	ID              int       `json:"id"`
	ContractAddress string    `json:"contract_address"`
	EventName       string    `json:"event_name"`
	EventData       string    `json:"event_data"` // JSONB stored as string
	BlockNumber     int64     `json:"block_number"`
	BlockHash       string    `json:"block_hash"`
	TransactionHash string    `json:"transaction_hash"`
	LogIndex        int       `json:"log_index"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
}

// Result represents tallied poll results
type Result struct {
	ID              int       `json:"id"`
	PollAddress     string    `json:"poll_address"`
	VoteCounts      []int     `json:"vote_counts"`
	TotalVotes      int       `json:"total_votes"`
	TalliedAt       time.Time `json:"tallied_at"`
	BlockNumber     int64     `json:"block_number"`
	TransactionHash string    `json:"transaction_hash"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
}
