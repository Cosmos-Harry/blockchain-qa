#!/bin/bash
set -e

echo "========================================="
echo "  Blockchain QA - Complete Setup"
echo "========================================="
echo ""

cd /Users/harrycosmos/go/projects/blockchain-qa

# Step 1: Check if Anvil is running
echo "Step 1: Checking if Anvil is running..."
if lsof -i :8545 > /dev/null 2>&1; then
    echo "✓ Anvil is already running on port 8545"
else
    echo "✗ Anvil is not running"
    echo ""
    echo "Please open a NEW terminal and run:"
    echo "  ~/.foundry/bin/anvil --port 8545 --accounts 20 --balance 10000"
    echo ""
    echo "Then come back here and run this script again."
    echo ""
    exit 1
fi

# Step 2: Deploy contracts
echo ""
echo "Step 2: Deploying contracts to Anvil..."
cd contracts
~/.foundry/bin/forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --broadcast

# Extract PollFactory address
FACTORY_ADDRESS=$(find broadcast/Deploy.s.sol -name "*.json" -type f -exec cat {} \; | jq -r '.transactions[] | select(.contractName == "PollFactory") | .contractAddress' | head -1)

if [ -z "$FACTORY_ADDRESS" ]; then
    echo "✗ Failed to extract PollFactory address"
    exit 1
fi

echo "✓ Contracts deployed!"
echo "  PollFactory: $FACTORY_ADDRESS"

# Step 3: Build CLI
echo ""
echo "Step 3: Building CLI (without CGO to avoid hang)..."
cd ../cli
CGO_ENABLED=0 go build -o poll-cli cmd/poll-cli/main.go
echo "✓ CLI built successfully"

# Step 4: Create .env file
echo ""
echo "Step 4: Creating .env file..."
cat > .env << EOF
RPC_URL=http://localhost:8545
POLL_FACTORY_ADDRESS=$FACTORY_ADDRESS
PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
EOF
echo "✓ .env file created"

# Step 5: Test CLI
echo ""
echo "Step 5: Testing CLI..."
export POLL_FACTORY_ADDRESS=$FACTORY_ADDRESS

echo ""
echo "Creating a test poll..."
./poll-cli create-poll \
  --private-key ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
  --question "What's your favorite Pokemon starter gen1?" \
  --options "Charmander,Squirtle,Bulbasaur" | grep -E "Poll Address|Transaction|Block Number"

echo ""
echo "========================================="
echo "  ✓ Setup Complete!"
echo "========================================="
echo ""
echo "You can now use the CLI:"
echo "  cd /Users/harrycosmos/go/projects/blockchain-qa/cli"
echo "  export POLL_FACTORY_ADDRESS=$FACTORY_ADDRESS"
echo "  ./poll-cli --help"
echo ""
