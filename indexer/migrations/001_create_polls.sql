-- Create polls table
CREATE TABLE IF NOT EXISTS polls (
    id SERIAL PRIMARY KEY,
    contract_address VARCHAR(42) NOT NULL UNIQUE,
    question TEXT NOT NULL,
    options TEXT[] NOT NULL,
    duration INTEGER NOT NULL,
    voter_merkle_root VARCHAR(66) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    closes_at TIMESTAMP NOT NULL,
    state VARCHAR(20) NOT NULL DEFAULT 'active',
    creator VARCHAR(42) NOT NULL,
    block_number BIGINT NOT NULL,
    transaction_hash VARCHAR(66) NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create votes table
CREATE TABLE IF NOT EXISTS votes (
    id SERIAL PRIMARY KEY,
    poll_address VARCHAR(42) NOT NULL,
    voter VARCHAR(42) NOT NULL,
    commitment VARCHAR(66) NOT NULL,
    choice INTEGER,
    nonce BYTEA,
    revealed BOOLEAN DEFAULT FALSE,
    committed_at TIMESTAMP NOT NULL,
    revealed_at TIMESTAMP,
    block_number BIGINT NOT NULL,
    transaction_hash VARCHAR(66) NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(poll_address, voter)
);

-- Create events table for raw event logs
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    contract_address VARCHAR(42) NOT NULL,
    event_name VARCHAR(50) NOT NULL,
    event_data JSONB NOT NULL,
    block_number BIGINT NOT NULL,
    block_hash VARCHAR(66) NOT NULL,
    transaction_hash VARCHAR(66) NOT NULL,
    log_index INTEGER NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(transaction_hash, log_index)
);

-- Create results table for tallied poll results
CREATE TABLE IF NOT EXISTS results (
    id SERIAL PRIMARY KEY,
    poll_address VARCHAR(42) NOT NULL UNIQUE,
    vote_counts INTEGER[] NOT NULL,
    total_votes INTEGER NOT NULL,
    tallied_at TIMESTAMP NOT NULL,
    block_number BIGINT NOT NULL,
    transaction_hash VARCHAR(66) NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for efficient queries
CREATE INDEX IF NOT EXISTS idx_polls_state ON polls(state);
CREATE INDEX IF NOT EXISTS idx_polls_created_at ON polls(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_polls_closes_at ON polls(closes_at);

CREATE INDEX IF NOT EXISTS idx_votes_poll_address ON votes(poll_address);
CREATE INDEX IF NOT EXISTS idx_votes_voter ON votes(voter);
CREATE INDEX IF NOT EXISTS idx_votes_revealed ON votes(poll_address, revealed);

CREATE INDEX IF NOT EXISTS idx_events_contract ON events(contract_address);
CREATE INDEX IF NOT EXISTS idx_events_name ON events(event_name);
CREATE INDEX IF NOT EXISTS idx_events_block ON events(block_number DESC);

CREATE INDEX IF NOT EXISTS idx_results_poll ON results(poll_address);
