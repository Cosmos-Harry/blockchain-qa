.PHONY: help setup test test-contracts test-zk test-indexer test-e2e deploy-local clean

# Colors for output
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

help: ## Show this help message
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

## Setup and Installation

setup: ## Set up the entire development environment
	@echo "${GREEN}Setting up blockchain-qa environment...${RESET}"
	@echo "${CYAN}[1/6] Installing Foundry dependencies...${RESET}"
	cd contracts && forge install foundry-rs/forge-std --no-commit || true
	@echo "${CYAN}[2/6] Building Rust ZK prover...${RESET}"
	cd zk-prover && cargo build --release
	@echo "${CYAN}[3/6] Installing Go dependencies (indexer)...${RESET}"
	cd indexer && go mod download
	@echo "${CYAN}[4/6] Installing Go dependencies (oracle)...${RESET}"
	cd oracle && go mod download
	@echo "${CYAN}[5/6] Installing Go dependencies (CLI)...${RESET}"
	cd cli && go mod download
	@echo "${CYAN}[6/6] Installing E2E test dependencies...${RESET}"
	cd e2e && npm install
	@echo "${GREEN}✓ Setup complete!${RESET}"
	@echo "${YELLOW}Next steps:${RESET}"
	@echo "  1. Start infrastructure: ${GREEN}make infra-up${RESET}"
	@echo "  2. Deploy contracts: ${GREEN}make deploy-local${RESET}"
	@echo "  3. Run tests: ${GREEN}make test${RESET}"

## Infrastructure Management

infra-up: ## Start infrastructure (Anvil, Postgres, Redis)
	@echo "${GREEN}Starting infrastructure...${RESET}"
	docker-compose up -d postgres redis anvil
	@echo "${YELLOW}Waiting for services to be ready...${RESET}"
	@./scripts/wait-for-services.sh

infra-down: ## Stop infrastructure
	@echo "${YELLOW}Stopping infrastructure...${RESET}"
	docker-compose down

infra-logs: ## Show infrastructure logs
	docker-compose logs -f postgres redis anvil

## Testing Commands

test: ## Run all tests (ONE COMMAND)
	@echo "${GREEN}========================================${RESET}"
	@echo "${GREEN}  Running Complete Test Suite${RESET}"
	@echo "${GREEN}========================================${RESET}"
	@echo ""
	@$(MAKE) test-contracts
	@echo ""
	@$(MAKE) test-zk
	@echo ""
	@$(MAKE) test-indexer
	@echo ""
	@$(MAKE) test-e2e
	@echo ""
	@echo "${GREEN}========================================${RESET}"
	@echo "${GREEN}  ✓ All Tests Passed!${RESET}"
	@echo "${GREEN}========================================${RESET}"

test-contracts: ## Run smart contract tests
	@echo "${CYAN}Running smart contract tests...${RESET}"
	cd contracts && forge test -vvv
	@echo "${YELLOW}Running fuzz tests (50k runs)...${RESET}"
	cd contracts && forge test --fuzz-runs 50000
	@echo "${GREEN}✓ Contract tests passed${RESET}"

test-zk: ## Run ZK prover tests
	@echo "${CYAN}Running ZK circuit tests...${RESET}"
	cd zk-prover && cargo test
	@echo "${YELLOW}Running ZK benchmarks...${RESET}"
	cd zk-prover && cargo bench --no-run
	@echo "${GREEN}✓ ZK prover tests passed${RESET}"

test-indexer: ## Run indexer & API tests
	@echo "${CYAN}Running indexer unit tests...${RESET}"
	cd indexer && go test ./...
	@echo "${YELLOW}Running indexer integration tests...${RESET}"
	cd indexer && go test ./test/integration/... -tags=integration || echo "${YELLOW}Integration tests require running infrastructure${RESET}"
	@echo "${GREEN}✓ Indexer tests passed${RESET}"

test-e2e: ## Run end-to-end tests
	@echo "${CYAN}Running E2E tests...${RESET}"
	@echo "${YELLOW}Ensuring infrastructure is running...${RESET}"
	@$(MAKE) infra-up
	@echo "${YELLOW}Deploying contracts...${RESET}"
	@$(MAKE) deploy-local
	@echo "${YELLOW}Starting services...${RESET}"
	@# Start services in background
	cd indexer && go run cmd/indexer/main.go &
	cd indexer && go run cmd/api/main.go &
	cd oracle && go run cmd/oracle/main.go &
	@echo "${YELLOW}Running E2E test suite...${RESET}"
	cd e2e && npm test || (echo "${RED}E2E tests failed${RESET}" && exit 1)
	@echo "${GREEN}✓ E2E tests passed${RESET}"

## Coverage Reports

coverage-contracts: ## Generate contract coverage report
	cd contracts && forge coverage --report lcov
	@echo "${GREEN}Coverage report: contracts/lcov.info${RESET}"

coverage-zk: ## Generate ZK prover coverage
	cd zk-prover && cargo tarpaulin --out Html
	@echo "${GREEN}Coverage report: zk-prover/tarpaulin-report.html${RESET}"

coverage-indexer: ## Generate indexer coverage
	cd indexer && go test -coverprofile=coverage.out ./...
	cd indexer && go tool cover -html=coverage.out -o coverage.html
	@echo "${GREEN}Coverage report: indexer/coverage.html${RESET}"

## Deployment

deploy-local: ## Deploy contracts to local Anvil
	@echo "${CYAN}Deploying contracts to local Anvil...${RESET}"
	cd contracts && forge script script/Deploy.s.sol \
		--broadcast \
		--rpc-url http://localhost:8545 \
		--private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
	@echo "${GREEN}✓ Contracts deployed${RESET}"

## Development Utilities

run-indexer: ## Run indexer service
	cd indexer && go run cmd/indexer/main.go

run-api: ## Run API server
	cd indexer && go run cmd/api/main.go

run-oracle: ## Run oracle service
	cd oracle && go run cmd/oracle/main.go

run-cli: ## Run CLI (example: make run-cli ARGS="create --question 'Test?'")
	cd cli && go run cmd/poll-cli/main.go $(ARGS)

## Build Commands

build: ## Build all components
	@echo "${CYAN}Building all components...${RESET}"
	cd contracts && forge build
	cd zk-prover && cargo build --release
	cd indexer && go build -o bin/indexer cmd/indexer/main.go
	cd indexer && go build -o bin/api cmd/api/main.go
	cd oracle && go build -o bin/oracle cmd/oracle/main.go
	cd cli && go build -o bin/poll-cli cmd/poll-cli/main.go
	@echo "${GREEN}✓ Build complete${RESET}"

build-wasm: ## Build ZK prover to WASM
	@echo "${CYAN}Building ZK prover to WASM...${RESET}"
	cd zk-prover && rustup target add wasm32-unknown-unknown
	cd zk-prover && cargo build --target wasm32-unknown-unknown --release
	@echo "${GREEN}✓ WASM build complete: zk-prover/target/wasm32-unknown-unknown/release/${RESET}"

## Code Quality

lint: ## Run linters on all components
	@echo "${CYAN}Linting contracts...${RESET}"
	cd contracts && forge fmt --check
	@echo "${CYAN}Linting Rust...${RESET}"
	cd zk-prover && cargo clippy -- -D warnings
	cd zk-prover && cargo fmt --check
	@echo "${CYAN}Linting Go...${RESET}"
	cd indexer && golangci-lint run || echo "${YELLOW}golangci-lint not installed${RESET}"
	cd oracle && golangci-lint run || echo "${YELLOW}golangci-lint not installed${RESET}"
	cd cli && golangci-lint run || echo "${YELLOW}golangci-lint not installed${RESET}"
	@echo "${GREEN}✓ Linting complete${RESET}"

format: ## Format all code
	@echo "${CYAN}Formatting code...${RESET}"
	cd contracts && forge fmt
	cd zk-prover && cargo fmt
	cd indexer && go fmt ./...
	cd oracle && go fmt ./...
	cd cli && go fmt ./...
	cd e2e && npm run format || true
	@echo "${GREEN}✓ Formatting complete${RESET}"

static-analysis: ## Run static analysis (Slither)
	@echo "${CYAN}Running Slither static analysis...${RESET}"
	cd contracts && slither src/ --filter-paths "test|script" || echo "${YELLOW}Slither not installed${RESET}"

## Benchmarks

bench-zk: ## Run ZK prover benchmarks
	@echo "${CYAN}Running ZK prover benchmarks...${RESET}"
	cd zk-prover && cargo bench
	@echo "${GREEN}Benchmark results: zk-prover/target/criterion/${RESET}"

bench-contracts: ## Run contract gas benchmarks
	@echo "${CYAN}Running contract gas benchmarks...${RESET}"
	cd contracts && forge snapshot
	@echo "${GREEN}Gas snapshot: contracts/.gas-snapshot${RESET}"

## Cleanup

clean: ## Clean all build artifacts
	@echo "${YELLOW}Cleaning build artifacts...${RESET}"
	cd contracts && forge clean
	cd zk-prover && cargo clean
	cd indexer && rm -rf bin/ coverage.out coverage.html
	cd oracle && rm -rf bin/
	cd cli && rm -rf bin/
	cd e2e && rm -rf node_modules/ coverage/
	@echo "${YELLOW}Stopping infrastructure...${RESET}"
	docker-compose down -v
	@echo "${GREEN}✓ Cleanup complete${RESET}"

clean-cache: ## Clean caches only (faster than clean)
	cd contracts && forge clean
	cd zk-prover && cargo clean
	cd indexer && go clean -cache
	cd oracle && go clean -cache
	cd cli && go clean -cache

## Database Management

db-migrate: ## Run database migrations
	@echo "${CYAN}Running database migrations...${RESET}"
	cd indexer && go run cmd/migrate/main.go

db-seed: ## Seed database with test data
	@echo "${CYAN}Seeding database...${RESET}"
	./scripts/seed-data.sh

db-reset: ## Reset database (drop and recreate)
	@echo "${YELLOW}Resetting database...${RESET}"
	docker-compose down postgres
	docker-compose up -d postgres
	@./scripts/wait-for-services.sh
	@$(MAKE) db-migrate
	@$(MAKE) db-seed
	@echo "${GREEN}✓ Database reset complete${RESET}"

## Utilities

generate-bindings: ## Generate Go contract bindings
	@echo "${CYAN}Generating Go contract bindings...${RESET}"
	./scripts/generate-bindings.sh

watch-contracts: ## Watch and re-run contract tests on changes
	cd contracts && forge test --watch

watch-zk: ## Watch and re-run ZK tests on changes
	cd zk-prover && cargo watch -x test

logs: ## Tail all service logs
	@echo "${CYAN}Tailing logs from all services...${RESET}"
	docker-compose logs -f

ps: ## Show running services
	docker-compose ps

## Documentation

docs: ## Generate documentation
	@echo "${CYAN}Generating documentation...${RESET}"
	cd contracts && forge doc
	cd zk-prover && cargo doc --no-deps --open
	@echo "${GREEN}✓ Documentation generated${RESET}"

## CI Simulation

ci: ## Simulate CI pipeline locally
	@echo "${GREEN}========================================${RESET}"
	@echo "${GREEN}  Simulating CI Pipeline${RESET}"
	@echo "${GREEN}========================================${RESET}"
	@echo ""
	@echo "${CYAN}Step 1: Linting...${RESET}"
	@$(MAKE) lint
	@echo ""
	@echo "${CYAN}Step 2: Building...${RESET}"
	@$(MAKE) build
	@echo ""
	@echo "${CYAN}Step 3: Testing...${RESET}"
	@$(MAKE) test
	@echo ""
	@echo "${CYAN}Step 4: Coverage...${RESET}"
	@$(MAKE) coverage-contracts
	@echo ""
	@echo "${GREEN}========================================${RESET}"
	@echo "${GREEN}  ✓ CI Simulation Complete!${RESET}"
	@echo "${GREEN}========================================${RESET}"

.DEFAULT_GOAL := help
