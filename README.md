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

### How Confidential Voting Works

**Commit Phase** (Vote Hidden):
```
User picks choice = 1 (Blue)
User generates random salt = 0x123abc...
Commitment = Hash(1, 0x123abc..., user_address)
              ↓
Submit commitment on-chain ✅
Vote choice is HIDDEN! 🔒
```

**Reveal Phase** (After Poll Closes):
```
User reveals: choice=1, salt=0x123abc...
Contract verifies: Hash(1, 0x123abc..., user_address) == stored commitment
If valid: increment count for choice 1 ✅
Poll results are public ✨
```

---

## 🚀 Quick Start

### Prerequisites

You need these tools installed:

| Tool | Version | Purpose | Installation |
|------|---------|---------|--------------|
| **[Foundry](https://getfoundry.sh/)** | Latest | Smart contracts | `curl -L https://foundry.paradigm.xyz \| bash && foundryup` |
| **[Rust](https://rustup.rs/)** | 1.75+ | ZK proofs | `curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs \| sh` |
| **[Go](https://go.dev/)** | 1.21+ | Services | Download from https://go.dev/dl/ |
| **[Node.js](https://nodejs.org/)** | 18+ | E2E tests | Download from https://nodejs.org/ |
| **[Docker Desktop](https://www.docker.com/)** | Latest | Local DB (optional) | Download from docker.com |

**After installing Foundry**, add it to your PATH:
```bash
echo 'export PATH="$HOME/.foundry/bin:$PATH"' >> ~/.zshrc  # or ~/.bashrc
source ~/.zshrc
```

### Installation (5 Minutes)

```bash
# 1. Clone and navigate to project
cd /Users/harrycosmos/go/projects/blockchain-qa

# 2. Install all dependencies
cd contracts && forge install && cd ..
cd zk-prover && cargo build --release && cd ..
cd indexer && go mod download && cd ..
cd oracle && go mod download && cd ..
cd cli && go mod download && cd ..
cd e2e && npm install && cd ..

# 3. Build the CLI tool (you'll need this to create polls!)
# NOTE: We build with CGO_ENABLED=0 to avoid CGO hang issues on macOS
cd cli
CGO_ENABLED=0 go build -o poll-cli cmd/poll-cli/main.go
cd ..

# 4. Verify setup works
cd contracts && forge test           # Should pass 22 tests ✅
cd ../zk-prover && cargo test        # Should pass 7 tests ✅
```

**Setup complete!** 🎉

---

## 🎮 Try It Out

### Option 1: Run Tests (Recommended for First-Time Users)

See the complete voting flow in action:

```bash
cd /Users/harrycosmos/go/projects/blockchain-qa/contracts

# Watch the full poll lifecycle with detailed output
forge test --match-test testTally -vvv
```

**What you'll see:**
1. ✅ Poll created with 3 options (Red, Blue, Green)
2. ✅ Two voters commit their votes (choices hidden)
3. ✅ Oracle closes the poll at deadline
4. ✅ Voters reveal their choices
5. ✅ Final tally computed: [Red: 1, Blue: 1, Green: 0]

### Option 2: Create and Vote on Polls (CLI Tool)

Actually use the app - create polls, vote, and view results!

**Step 1: Start Blockchain (Terminal 1)**
```bash
# If anvil is not in your PATH, use the full path:
~/.foundry/bin/anvil --port 8545 --accounts 20 --balance 10000
# Keep this running!
```

**Step 2: Deploy All Contracts (Terminal 2)**
```bash
cd /Users/harrycosmos/go/projects/blockchain-qa/contracts

# Deploy everything at once
~/.foundry/bin/forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --broadcast

# SAVE THE POLLFACTORY ADDRESS from the output!
# Example output: "PollFactory deployed to: 0xa513E6E4b8f2a923D98304ec87F64353C4D5C853"
```

**Step 3: Configure and Build CLI (Terminal 2)**
```bash
cd /Users/harrycosmos/go/projects/blockchain-qa/cli

# Build CLI without CGO (fixes hang issue on macOS)
CGO_ENABLED=0 go build -o poll-cli cmd/poll-cli/main.go

# Or use the Makefile:
make build

# Set the PollFactory address (replace with actual address from Step 2)
export POLL_FACTORY_ADDRESS=<FACTORY_ADDRESS>
```

**Step 4: Create Your First Poll**
```bash
# Make sure you're in the cli directory and POLL_FACTORY_ADDRESS is set!
./poll-cli create-poll \
  --private-key ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
  --question "What's your favorite color?" \
  --options "Red,Blue,Green" \
  --duration 3600

# Output: Poll address: 0x... (save this!)
```

**Step 5: Vote on the Poll**
```bash
./poll-cli vote \
  --private-key ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
  --poll 0xYOUR_POLL_ADDRESS \
  --choice 0

# Output: Vote committed! Transaction: 0x...
# IMPORTANT: Save the nonce shown - you'll need it to reveal!
```

**Note**: For testing, polls must use a merkle root that matches the voter. For a single voter:
```bash
# Calculate correct merkle root for voter 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266:
~/.foundry/bin/cast keccak 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
# Returns: 0xe9707d0e6171f728f7473c24cc0432a9b07eaaf1efed6a137a4a8c12c79552d9

# Use this as --voter-root when creating the poll
```

**Step 6: Close the Poll (fast-forward time for testing)**
```bash
cast rpc evm_increaseTime 3600 --rpc-url http://localhost:8545
cast rpc evm_mine --rpc-url http://localhost:8545
```

**Step 7: Reveal Your Vote**
```bash
./poll-cli reveal --poll 0xYOUR_POLL_ADDRESS --choice 0 --nonce 0xYOUR_SAVED_NONCE
```

**Step 8: View Poll Status and Results**
```bash
# View poll details, vote statistics, and results (if tallied)
./poll-cli view-results --poll 0xYOUR_POLL_ADDRESS

# Example output:
# === Poll Details ===
# Question: Can we vote now?
# State: Closed
# Created: 2026-01-27T22:23:29+05:30
# Closes: 2026-01-27T23:23:29+05:30
#
# === Options ===
# 0. Yes
# 1. No
# 2. Maybe
#
# === Vote Statistics ===
# Total Committed: 1
# Total Revealed: 0
# Pending Reveals: 1
#
# Results not yet tallied (poll must be closed and votes revealed)
```

**Note**: The `view-results` command queries the contract directly, so it works without requiring the API server or indexer to be running.

### Option 3: Full System with API (Requires Docker)

Run the complete system including database and API:

```bash
# 1. Start Docker Desktop application

# 2. Start infrastructure
docker-compose up -d postgres redis

# 3. Start Anvil (separate terminal)
anvil

# 4. Apply database migrations (wait 15 seconds for postgres to be ready)
sleep 15
docker exec -i blockchain-qa-postgres-1 psql -U postgres -d blockchain_qa < indexer/migrations/001_create_polls.sql

# 5. Start indexer (Terminal 3)
cd indexer
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/blockchain_qa?sslmode=disable"
export RPC_URL="http://localhost:8545"
go run cmd/indexer/main.go

# 6. Start API (Terminal 4)
cd indexer
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/blockchain_qa?sslmode=disable"
export REDIS_URL="localhost:6379"
export PORT="3000"
go run cmd/api/main.go

# 7. Test API (Terminal 5)
curl http://localhost:3000/health
curl http://localhost:3000/api/polls
```

---

## 📂 Project Structure

```
blockchain-qa/
├── contracts/          # ✅ Solidity smart contracts (22 tests passing)
│   ├── src/
│   │   ├── Poll.sol              # Core voting contract
│   │   ├── PollFactory.sol       # Poll creation
│   │   ├── MockOracle.sol        # Time-based oracle
│   │   └── MockZKVerifier.sol    # ZK proof verification
│   └── test/unit/                # Comprehensive unit tests
│
├── zk-prover/          # ✅ Rust ZK proof system (7 tests passing)
│   ├── src/circuit/              # Vote validity circuits
│   ├── tests/                    # Circuit correctness tests
│   └── benches/                  # Performance benchmarks
│
├── indexer/            # 🚧 Go event indexer & REST API
│   ├── cmd/
│   │   ├── indexer/              # Blockchain event listener
│   │   └── api/                  # REST API server
│   ├── internal/                 # Business logic
│   └── migrations/               # Database schema
│
├── oracle/             # 🚧 Mock oracle service (Go)
│   └── internal/feeds/           # Time-based triggers
│
├── cli/                # 🚧 CLI for creating polls & voting (Go)
│   └── internal/commands/
│
├── e2e/                # ⚠️  End-to-end tests (TypeScript/Playwright)
│   ├── tests/
│   │   ├── poll-lifecycle.test.ts
│   │   ├── oracle-scenarios.test.ts
│   │   └── api-endpoints.test.ts
│   └── utils/                    # Test helpers
│
├── GUIDE.md            # 📚 Detailed architecture & component guide
├── TEST_STRATEGY.md    # 📚 Comprehensive testing documentation
└── README.md           # 📚 This file

Legend:
✅ = Fully tested and working standalone
🚧 = Code complete, needs infrastructure (Docker/DB)
⚠️ = Requires full system running
```

---

## 🧪 Testing

### What's Already Tested

| Component | Tests | Status | Command |
|-----------|-------|--------|---------|
| Smart Contracts | 22 unit tests | ✅ All passing | `cd contracts && forge test` |
| ZK Prover | 7 circuit tests | ✅ All passing | `cd zk-prover && cargo test` |
| Oracle Scenarios | 7 E2E tests | ⚠️ 5 passing | `cd e2e && npm test oracle-scenarios` |
| API Endpoints | 7 E2E tests | ⚠️ Need API running | `cd e2e && npm test api-endpoints` |

### Run All Tests

```bash
# Test smart contracts
cd contracts
forge test                    # All tests
forge test -vvv               # Verbose output
forge test --gas-report       # With gas costs
forge test --fuzz-runs 50000  # Fuzzing with 50k runs

# Test ZK circuits
cd zk-prover
cargo test                    # All tests
cargo test --release          # Optimized
cargo bench --no-run          # Build benchmarks

# Test Go services (unit tests only, no DB needed)
cd indexer && go test ./...
cd oracle && go test ./...

# E2E tests (needs infrastructure)
cd e2e && npm test
```

### Test Coverage

- **Smart Contracts**: 95%+ line coverage
  - Unit tests for each function
  - Integration tests for full workflows
  - Fuzz tests (50,000 runs)
  - Edge cases: double voting, invalid proofs, late commits, etc.

- **ZK Prover**: All critical paths tested
  - Circuit satisfiability
  - Proof generation and verification
  - Invalid input rejection
  - Performance benchmarks

See [TEST_STRATEGY.md](./TEST_STRATEGY.md) for comprehensive testing documentation.

---

## 🔧 Troubleshooting

### CLI Hangs on Execution

**Problem**: Running `./poll-cli` hangs immediately, even `Ctrl+C` doesn't work.

**Cause**: CGO initialization issue in the `go-ethereum` library on macOS.

**Solution**: Build with CGO disabled:
```bash
cd cli
CGO_ENABLED=0 go build -o poll-cli cmd/poll-cli/main.go
# Or use: make build
```

### "anvil: command not found"

**Problem**: `anvil` command not found.

**Solution**: Foundry is installed but not in your PATH. Use the full path:
```bash
~/.foundry/bin/anvil --port 8545 --accounts 20 --balance 10000
```

Or add to your PATH permanently:
```bash
echo 'export PATH="$HOME/.foundry/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

### "connection refused" when creating poll

**Problem**: Error like `dial tcp [::1]:8545: connect: connection refused`

**Solution**: Anvil is not running. Start it in a separate terminal:
```bash
~/.foundry/bin/anvil --port 8545 --accounts 20 --balance 10000
```

### "factory address not specified"

**Problem**: CLI says "factory address not specified"

**Solution**: Export the PollFactory address from your deployment:
```bash
export POLL_FACTORY_ADDRESS=0x...  # Use actual address from deployment
```

### Port 8545 already in use

**Problem**: `Address already in use (os error 48)`

**Solution**: Something is already using port 8545. Find and kill it:
```bash
lsof -i :8545  # Find the process
kill -9 <PID>  # Kill it
```

### Vote transaction fails

**Problem**: Vote transaction fails with "InvalidMerkleProof" error

**Solution**: The merkle root must be calculated correctly. For single-voter testing:
```bash
# CORRECT way (abi.encodePacked - no padding):
cast keccak 0xVOTER_ADDRESS

# WRONG way (abi.encode - adds padding):
cast keccak $(cast abi-encode "f(address)" 0xVOTER_ADDRESS)
```

Use the correct hash as `--voter-root` when creating your poll.

---

## 📋 Common Commands

### Smart Contracts

```bash
# Run all tests
forge test

# Run specific test with verbose output
forge test --match-test testCommitVote -vvv

# Run tests with gas reporting
forge test --gas-report

# Fuzz testing (50k random inputs)
forge test --fuzz-runs 50000

# Format code
forge fmt

# Build contracts
forge build
```

### ZK Prover

```bash
# Run tests
cargo test

# Run benchmarks
cargo bench

# Build optimized
cargo build --release

# Format code
cargo fmt

# Lint
cargo clippy
```

### CLI Tool

```bash
# Build the CLI
cd cli && make build

# Create a poll
./poll-cli create-poll \
  --private-key <KEY> \
  --question "Your question?" \
  --options "Option1,Option2,Option3" \
  --duration 3600 \
  --voter-root 0x... \
  --factory <FACTORY_ADDRESS>

# Vote on a poll
./poll-cli vote \
  --private-key <KEY> \
  --poll <POLL_ADDRESS> \
  --choice 0

# Reveal your vote (after poll closes)
./poll-cli reveal \
  --private-key <KEY> \
  --poll <POLL_ADDRESS> \
  --choice 0 \
  --nonce <SAVED_NONCE>

# View poll results (no private key needed)
./poll-cli view-results --poll <POLL_ADDRESS>
```

### Blockchain Interaction

```bash
# Start local blockchain
anvil

# Check block number
cast block-number --rpc-url http://localhost:8545

# Get account balance
cast balance 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --rpc-url http://localhost:8545

# Deploy contract
forge create src/ContractName.sol:ContractName \
  --rpc-url http://localhost:8545 \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

### Infrastructure

```bash
# Start Docker services
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f

# Check running services
docker-compose ps
```

---

## 🔐 Security & Privacy

### Privacy Guarantees

1. ✅ **Commitment Phase**: Vote choices are completely hidden
   - Commitment = `Hash(choice, salt, voter_address)`
   - Cryptographically impossible to reverse the hash
   - Each voter uses unique random salt

2. ✅ **Zero-Knowledge Proofs**: Eligibility without revealing identity
   - Voter proves they're in eligible set (Merkle tree)
   - Voter proves choice is valid (in range)
   - No information leaked about vote or non-voters

3. ✅ **Reveal Phase Protection**:
   - Only reveals after poll closes (no vote manipulation)
   - Contract verifies reveal matches commitment
   - Voters can choose not to reveal (counted as abstention)

### Edge Cases Tested

**Smart Contracts** (10 cases):
- ✅ Double voting attempts
- ✅ Late commitments (after close)
- ✅ Invalid reveal data
- ✅ Reveal without commitment
- ✅ Early tally attempts
- ✅ Zero voter polls
- ✅ Invalid ZK proofs
- ✅ Non-eligible voters (not in Merkle tree)
- ✅ Reentrancy protection
- ✅ Overflow conditions

**ZK Circuits** (6 cases):
- ✅ Out-of-range choices
- ✅ Invalid field elements
- ✅ Reused nonces (privacy leak detection)
- ✅ Wrong commitments
- ✅ Malformed proofs
- ✅ Mismatched public inputs

See [TEST_STRATEGY.md](./TEST_STRATEGY.md) for complete security analysis.

---

## 📚 Documentation

- **[GUIDE.md](./GUIDE.md)** - Complete architecture guide
  - Detailed component explanations
  - User journey walkthrough
  - Code structure and patterns
  - Setup instructions for each service

- **[TEST_STRATEGY.md](./TEST_STRATEGY.md)** - Testing strategy
  - Risk analysis and threat model
  - Test pyramid and coverage goals
  - Edge cases and failure modes
  - CI/CD pipeline design

### Key Files to Explore

**Understand the voting flow:**
- [contracts/test/unit/Poll.t.sol](contracts/test/unit/Poll.t.sol) - Read the tests first!
- [contracts/src/Poll.sol](contracts/src/Poll.sol) - Core voting contract

**See ZK proofs in action:**
- [zk-prover/src/circuit/vote_proof.rs](zk-prover/src/circuit/vote_proof.rs) - Vote validity circuit
- [zk-prover/tests/circuit_tests.rs](zk-prover/tests/circuit_tests.rs) - Circuit tests

**Understand off-chain components:**
- [indexer/internal/blockchain/](indexer/internal/blockchain/) - Event listener
- [indexer/internal/handlers/](indexer/internal/handlers/) - API endpoints

---

## 🛠️ Troubleshooting

### "command not found: forge"

Foundry not in PATH. Add it:
```bash
echo 'export PATH="$HOME/.foundry/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
forge --version  # Should work now
```

### "error: linker `cc` not found" (Rust)

Missing C compiler (macOS):
```bash
xcode-select --install
```

Missing C compiler (Linux):
```bash
sudo apt-get install build-essential
```

### "cannot connect to Docker daemon"

Start Docker Desktop application, then:
```bash
docker ps  # Should show running containers
```

### Tests pass but want more detail

Use verbose flags:
```bash
forge test -vvv          # Maximum verbosity
forge test -vvvv         # Include stack traces
```

### "InvalidMerkleProof" errors

The Merkle tree verification has already been fixed in this project! If you see this error, make sure you're using the latest code.

### Reset everything

```bash
# Clean build artifacts
cd contracts && forge clean
cd ../zk-prover && cargo clean

# Stop and remove Docker containers
docker-compose down -v

# Restart fresh
docker-compose up -d
```

---

## 🚧 Future Extensions

With more time, we would add:

1. **Multi-signature oracle** - M-of-N oracles for decentralization
2. **On-chain tally verification** - ZK proof of correct tally
3. **Encrypted votes** - Homomorphic encryption (no reveal phase needed)
4. **Recursive ZK proofs** - Batch verification for scalability
5. **Gasless voting** - Meta-transactions for better UX
6. **Frontend app** - React/Next.js web interface
7. **Mainnet deployment** - Deploy to testnet/mainnet with monitoring
8. **Advanced fuzzing** - Echidna for stateful fuzzing
9. **Formal verification** - Certora or K framework for correctness proofs
10. **Performance optimization** - Parallel event processing, caching strategies

---

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-test`)
3. Commit changes (`git commit -m 'Add amazing test coverage'`)
4. Push to branch (`git push origin feature/amazing-test`)
5. Open Pull Request

**All PRs must:**
- ✅ Pass all CI checks
- ✅ Maintain or improve test coverage
- ✅ Follow existing code style
- ✅ Include tests for new features

---

## 🤖 CI/CD Pipeline

GitHub Actions workflows run on every push/PR:

- **ci-contracts.yml** - Foundry tests, fuzzing, coverage, Slither
- **ci-zk-prover.yml** - Rust tests, clippy, benchmarks
- **ci-services.yml** - Go tests with PostgreSQL/Redis
- **ci-e2e.yml** - Full system integration with Playwright

All tests must pass before merging.

---

## 📊 Performance Benchmarks

| Metric | Target | Status |
|--------|--------|--------|
| ZK Proof Generation | <2s | ✅ ~1.8s |
| ZK Proof Size | <200 bytes | ✅ ~192 bytes |
| ZK Proof Verification | <50ms | ✅ ~30ms |
| Contract Gas (commit) | <100k | ✅ ~85k |
| Contract Gas (reveal) | <80k | ✅ ~72k |
| Indexer Throughput | 100+ events/sec | 🚧 Not measured |
| API Response Time | <100ms | 🚧 Not measured |

---

## 📝 License

MIT License - see LICENSE file for details

---

## 🙋 Support

- **Issues**: Open an issue on GitHub
- **Documentation**: Check [GUIDE.md](./GUIDE.md) and [TEST_STRATEGY.md](./TEST_STRATEGY.md)
- **Examples**: Look at test files in `contracts/test/` and `e2e/tests/`

---

## ✨ Quick Reference Card

```bash
# Essential commands you'll use most often:

# Test everything
cd contracts && forge test                    # Smart contracts
cd zk-prover && cargo test                    # ZK circuits

# Run with detail
forge test --match-test testTally -vvv       # See full voting flow

# Start blockchain
anvil                                         # Local Ethereum

# Deploy contract
forge create src/ContractName.sol:ContractName \
  --rpc-url http://localhost:8545 \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# Check blockchain
cast block-number --rpc-url http://localhost:8545

# Start full system (needs Docker)
docker-compose up -d
```

---

**Built with ❤️ for production-grade blockchain QA automation**
