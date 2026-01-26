# End-to-End Tests

Comprehensive E2E test suite using Playwright and TypeScript to verify the complete poll dApp workflow.

## Test Coverage

### Poll Lifecycle (`poll-lifecycle.test.ts`)
- Complete flow: create → commit → close → reveal → tally
- Vote privacy verification (commitments don't leak choice)
- API integration with indexer
- Blockchain state management (snapshots/reverts)
- Time manipulation for testing deadlines

### API Endpoints (`api-endpoints.test.ts`)
- Health check endpoint
- Poll listing with pagination and state filters
- Individual poll queries
- Vote statistics
- Results retrieval
- Error handling for non-existent resources
- Response time validation

### Oracle Scenarios (`oracle-scenarios.test.ts`)
- **On-time response**: Poll closes at exact deadline
- **Late response**: Poll closes 5-15 minutes late
- **Invalid data**: Oracle sends malformed transactions
- **No response**: Oracle downtime simulation
- Runtime mode switching
- Concurrent poll closures
- Gas price spike handling

## Setup

1. Install dependencies:
```bash
npm install
npx playwright install
```

2. Configure environment:
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. Start required services:
```bash
# Terminal 1: Start Anvil (local Ethereum node)
anvil

# Terminal 2: Start PostgreSQL and Redis
docker-compose up postgres redis

# Terminal 3: Start Indexer
cd ../indexer && go run cmd/indexer/main.go

# Terminal 4: Start API
cd ../indexer && go run cmd/api/main.go
```

## Running Tests

```bash
# Run all tests
npm test

# Run with UI mode (interactive)
npm run test:ui

# Run in headed mode (see browser)
npm run test:headed

# Run specific test file
npx playwright test tests/poll-lifecycle.test.ts

# Debug mode
npm run test:debug

# View test report
npm run report
```

## Test Structure

```
e2e/
├── tests/                    # Test files
│   ├── poll-lifecycle.test.ts
│   ├── api-endpoints.test.ts
│   └── oracle-scenarios.test.ts
├── utils/                    # Helper utilities
│   ├── blockchain.ts         # Blockchain interactions
│   └── api.ts               # API client
├── playwright.config.ts      # Playwright configuration
├── tsconfig.json            # TypeScript configuration
└── package.json             # Dependencies and scripts
```

## Key Features

- **Blockchain State Management**: Each test uses snapshots to ensure clean state
- **API Integration**: Tests verify indexer correctly processes on-chain events
- **Time Manipulation**: Use `evm_increaseTime` and `evm_mine` for deadline testing
- **Privacy Verification**: Ensure vote commitments maintain privacy
- **Oracle Testing**: Comprehensive coverage of all oracle response modes
- **Error Handling**: Tests verify proper error responses

## Notes

- Tests run sequentially (`workers: 1`) to avoid blockchain state conflicts
- Each test has 2-minute timeout for slow blockchain operations
- Snapshots are taken before each test and reverted after
- Uses Anvil's test accounts with pre-funded ETH
- Real contract deployment will be added once contract bindings are generated

## Future Enhancements

- Add contract deployment scripts
- Integrate with Foundry for contract interactions
- Add ZK proof generation tests (via WASM prover)
- Performance/load testing with many concurrent voters
- Reorg handling tests for indexer
- Full integration with deployed contracts
