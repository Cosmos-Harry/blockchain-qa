# Blockchain QA Automation - Confidential Poll dApp

A comprehensive QA automation framework demonstrating production-grade testing practices for blockchain systems, featuring smart contracts, off-chain services, ZK proof privacy, and complete test coverage.

## 🎯 Project Overview

This project implements a **confidential voting dApp** with:
- ✅ **Smart contracts** (Solidity) - Privacy-preserving voting with commit-reveal + ZK proofs
- ✅ **Off-chain indexer/API** (Go) - Event processing and query optimization
- ✅ **ZK proof system** (Rust/arkworks) - Vote privacy using Groth16 proofs
- ✅ **Mock oracle** - Configurable scenarios for testing edge cases
- ✅ **CLI** - User simulation tool for multi-account workflows
- ✅ **Comprehensive test suite** - Unit, integration, fuzzing, and E2E tests
- ✅ **CI/CD pipeline** - Automated testing on every commit

## 🏗️ Architecture

```
┌─────────────┐      ┌──────────────┐      ┌─────────────┐
│   Voters    │─────▶│ Poll Contract│◀─────│  Oracle     │
│   (CLI)     │      │  (Solidity)  │      │   (Go)      │
└─────────────┘      └──────────────┘      └─────────────┘
       │                     │
       │ ZK Proof            │ Events
       ▼                     ▼
┌─────────────┐      ┌──────────────┐
│ ZK Prover   │      │  Indexer     │──────▶ PostgreSQL
│  (Rust)     │      │   (Go)       │       Redis Cache
└─────────────┘      └──────────────┘
                            │
                            ▼
                     ┌──────────────┐
                     │  REST API    │
                     │   (Fiber)    │
                     └──────────────┘
```

### Core Components

1. **Smart Contracts** (`contracts/`)
   - `Poll.sol` - Core voting contract with commit-reveal mechanism
   - `PollFactory.sol` - Creates and manages polls
   - `MockOracle.sol` - Configurable oracle for testing (OnTime/Late/Invalid/NoResponse modes)
   - `ZKVerifier.sol` - Groth16 proof verification (generated from Rust circuit)

2. **ZK Proof System** (`zk-prover/`)
   - `VoteProofCircuit` - Proves vote validity (choice in range, commitment correctness)
   - `EligibilityProofCircuit` - Proves voter is in Merkle tree
   - Groth16 proving using arkworks (ark-groth16, ark-bn254)
   - WASM compilation for CLI/browser usage

3. **Off-Chain Indexer & API** (`indexer/`)
   - Event listener consuming contract events
   - PostgreSQL database for indexed data
   - Redis caching for performance
   - REST API (Fiber framework) for queries and submissions

4. **Mock Oracle** (`oracle/`)
   - Time-based feed triggering poll closing
   - Configurable response modes for testing
   - Simulates edge cases: late responses, invalid data, downtime

5. **CLI** (`cli/`)
   - Commands: `create-poll`, `vote`, `reveal`, `view-results`
   - Automatic ZK proof generation
   - Multi-account wallet management

## 🚀 Quick Start

### Prerequisites

- [Foundry](https://getfoundry.sh/) - Smart contract development
- [Rust](https://rustup.rs/) 1.75+ - ZK prover
- [Go](https://go.dev/) 1.21+ - Indexer, oracle, CLI
- [Node.js](https://nodejs.org/) 18+ - E2E tests
- [Docker](https://www.docker.com/) & Docker Compose - Local environment

### Installation

```bash
# Clone repository
git clone <repo-url>
cd blockchain-qa

# Run setup (installs dependencies, starts infrastructure)
make setup
```

### ⚡ One Command to Run All Tests

```bash
make test
```

This runs:
1. Contract tests (Foundry) - unit, integration, fuzzing
2. ZK prover tests (Rust) - circuit tests, benchmarks
3. Indexer tests (Go) - unit, integration tests
4. E2E tests (TypeScript) - full system integration

### Local Development

```bash
# Terminal 1: Start infrastructure (Anvil, Postgres, Redis)
docker-compose up

# Terminal 2: Deploy contracts
make deploy-local

# Terminal 3: Start indexer
cd indexer && make run-indexer

# Terminal 4: Start API
cd indexer && make run-api

# Terminal 5: Start oracle
cd oracle && go run cmd/oracle/main.go

# Terminal 6: Use CLI
cd cli
./poll-cli create --question "Favorite color?" --options "Red,Blue,Green" --duration 3600
./poll-cli vote --poll 0x... --choice 1
./poll-cli reveal --poll 0x...
./poll-cli results --poll 0x...
```

## 📋 Testing Strategy

See [TEST_STRATEGY.md](./TEST_STRATEGY.md) for comprehensive testing documentation.

### Test Coverage

- **Smart Contracts**: 95%+ line coverage
  - Unit tests for each contract function
  - Integration tests for full voting flow
  - Fuzz tests (50,000 runs) for edge cases
  - Static analysis with Slither

- **ZK Prover**: All circuits tested
  - Circuit satisfiability tests
  - Proof generation and verification
  - Performance benchmarks (<2s proof generation)

- **Indexer**: 80%+ coverage
  - Unit tests with mocked dependencies
  - Integration tests with real database
  - Reorg handling and error recovery

- **End-to-End**: Critical paths covered (Playwright + TypeScript)
  - Full poll lifecycle (create → commit → close → reveal → tally)
  - Privacy verification (zero-knowledge guarantees)
  - Oracle edge cases (late, invalid, no-response)
  - API endpoint testing with error handling
  - Blockchain time manipulation and state snapshots

### Edge Cases Tested

**Smart Contracts** (10 cases)
- ✅ Double voting attempts
- ✅ Late commitments after close
- ✅ Invalid reveal data
- ✅ Reveal without commitment
- ✅ Early tally attempts
- ✅ Zero voter polls
- ✅ Invalid ZK proofs
- ✅ Non-eligible voters
- ✅ Reentrancy attacks
- ✅ Overflow conditions

**ZK Circuits** (6 cases)
- ✅ Out-of-range choices
- ✅ Negative field elements
- ✅ Reused nonces
- ✅ Wrong commitments
- ✅ Malformed proofs
- ✅ Mismatched public inputs

**Indexer** (6 cases)
- ✅ Blockchain reorgs
- ✅ Missed events recovery
- ✅ Duplicate event handling
- ✅ Out-of-order events
- ✅ Database failures
- ✅ Invalid event data

**Oracle** (4 cases)
- ✅ Concurrent close requests
- ✅ Frontrunning attempts
- ✅ Gas price spikes
- ✅ Oracle downtime

## 🔧 Development Commands

```bash
# Setup environment
make setup              # One-time setup

# Testing
make test               # Run all tests
make test-contracts     # Smart contract tests only
make test-zk            # ZK prover tests only
make test-indexer       # Indexer tests only
make test-e2e           # End-to-end tests only

# Development
make deploy-local       # Deploy to local Anvil
make clean              # Clean build artifacts

# Individual component tests
cd contracts && forge test --match-contract PollTest
cd zk-prover && cargo test vote_proof
cd indexer && go test ./internal/handlers/...
cd e2e && npm test -- poll-lifecycle
```

## 📂 Project Structure

```
blockchain-qa/
├── contracts/          # Solidity smart contracts
│   ├── src/           # Contract source code
│   ├── test/          # Unit, integration, fuzz tests
│   └── script/        # Deployment scripts
├── zk-prover/         # Rust ZK proof system
│   ├── src/circuit/   # Vote and eligibility circuits
│   ├── tests/         # Circuit tests
│   └── benches/       # Performance benchmarks
├── indexer/           # Go event indexer & API
│   ├── cmd/           # Indexer and API entry points
│   ├── internal/      # Business logic
│   ├── migrations/    # Database migrations
│   └── test/          # Unit and integration tests
├── oracle/            # Mock oracle service
│   ├── cmd/           # Oracle entry point
│   └── internal/      # Feed logic
├── cli/               # User simulation CLI
│   └── internal/      # Commands and wallet
├── e2e/               # End-to-end tests
│   └── tests/         # Full system integration tests
├── scripts/           # Automation scripts
├── .github/workflows/ # CI/CD pipelines
├── docker-compose.yml # Local development environment
├── Makefile           # Unified build/test commands
├── README.md          # This file
└── TEST_STRATEGY.md   # Comprehensive test documentation
```

## 🔐 Privacy Model

The confidential poll uses a **commit-reveal scheme** with **zero-knowledge proofs**:

1. **Commit Phase**:
   - Voter generates: `commitment = Hash(choice, salt, voter_address)`
   - Voter proves eligibility with ZK proof (in Merkle tree)
   - Voter proves vote validity with ZK proof (choice in range)
   - Commitment submitted on-chain (vote choice hidden)

2. **Reveal Phase** (after poll closes):
   - Voter reveals `(choice, salt)`
   - Contract verifies: `Hash(choice, salt, voter) == commitment`
   - Vote counted in tally

3. **Privacy Guarantees**:
   - ✅ Commitments reveal no information about vote choice
   - ✅ ZK proofs are zero-knowledge (simul ator indistinguishable)
   - ✅ Only revealed votes are public
   - ✅ Merkle tree hides non-voter identities

## 🤖 CI/CD Pipeline

GitHub Actions workflows run on every push/PR:

- **ci-contracts.yml** - Foundry tests, fuzzing, coverage, Slither
- **ci-zk-prover.yml** - Rust tests, clippy, benchmarks
- **ci-services.yml** - Go tests for indexer, oracle, CLI with PostgreSQL/Redis
- **ci-e2e.yml** - Full system integration tests with Playwright, Anvil, all services

All tests must pass before merging. Logs and artifacts uploaded on failure.

## 📊 Performance Benchmarks

- ZK Proof Generation: <2s (target)
- ZK Proof Size: <200 bytes
- ZK Proof Verification: <50ms
- Indexer Event Processing: 100+ events/sec
- API Response Time: <100ms (with Redis cache)

## 🚧 Future Extensions

With more time, we would add:

1. **Multi-signature oracle** - M-of-N oracles (decentralization)
2. **On-chain tally verification** - ZK proof of correct tally
3. **Encrypted votes** - Homomorphic encryption (no reveal phase)
4. **Recursive ZK proofs** - Batch verification (scalability)
5. **Gasless voting** - Meta-transactions (UX improvement)
6. **Frontend app** - React/Next.js web interface
7. **Mainnet deployment** - Testnet/mainnet with monitoring
8. **Advanced fuzzing** - Echidna stateful fuzzing
9. **Formal verification** - Certora or K framework proofs
10. **Performance optimization** - Parallel processing, batching

## 📝 License

MIT License - see LICENSE file for details

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-test`)
3. Commit changes (`git commit -m 'Add amazing test coverage'`)
4. Push to branch (`git push origin feature/amazing-test`)
5. Open Pull Request

All PRs must:
- Pass all CI checks
- Maintain or improve test coverage
- Follow existing code style
- Include tests for new features

## 📚 Documentation

- [TEST_STRATEGY.md](./TEST_STRATEGY.md) - Comprehensive testing strategy
- [contracts/README.md](./contracts/README.md) - Smart contract documentation
- [zk-prover/README.md](./zk-prover/README.md) - ZK circuit documentation
- [indexer/README.md](./indexer/README.md) - API documentation

## 🙋 Support

For questions or issues:
- Open an issue on GitHub
- See documentation in `docs/`
- Check CI logs for test failures

---

**Built with**  for production-grade blockchain QA automation.
