package commands

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Cosmos-Harry/blockchain-qa/cli/internal/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var (
	closePollAddress string
	oracleAddress    string
)

var closePollCmd = &cobra.Command{
	Use:   "close-poll",
	Short: "Close a poll via the oracle",
	Long:  `Trigger the MockOracle to close a poll after its duration has passed.`,
	RunE:  runClosePoll,
}

func init() {
	rootCmd.AddCommand(closePollCmd)

	closePollCmd.Flags().StringVar(&closePollAddress, "poll", "", "Poll contract address")
	closePollCmd.Flags().StringVar(&oracleAddress, "oracle", "", "MockOracle contract address")

	mustMarkRequired(closePollCmd, "poll")
	mustMarkRequired(closePollCmd, "oracle")
}

// MockOracle ABI (only fulfillRequest method needed)
const mockOracleABI = `[{"inputs":[{"internalType":"address","name":"poll","type":"address"}],"name":"fulfillRequest","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

func runClosePoll(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	w, err := wallet.NewWallet(ctx, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	defer w.Close()

	log.Printf("Closing poll %s via oracle %s\n", closePollAddress, oracleAddress)

	oracleAddr := common.HexToAddress(oracleAddress)
	pollAddr := common.HexToAddress(closePollAddress)

	// Parse oracle ABI and pack fulfillRequest call
	oracleABI, err := abi.JSON(strings.NewReader(mockOracleABI))
	if err != nil {
		return fmt.Errorf("failed to parse oracle ABI: %w", err)
	}

	data, err := oracleABI.Pack("fulfillRequest", pollAddr)
	if err != nil {
		return fmt.Errorf("failed to pack fulfillRequest: %w", err)
	}

	// Verify selector matches expected
	expectedSig := crypto.Keccak256([]byte("fulfillRequest(address)"))[:4]
	if fmt.Sprintf("%x", data[:4]) != fmt.Sprintf("%x", expectedSig) {
		return fmt.Errorf("ABI pack produced unexpected selector")
	}

	auth, err := w.GetAuth(ctx)
	if err != nil {
		return fmt.Errorf("failed to get auth: %w", err)
	}

	// Send transaction directly using the packed data
	tx, err := w.SendRawTx(ctx, auth, oracleAddr, data)
	if err != nil {
		return fmt.Errorf("failed to send fulfillRequest: %w", err)
	}

	log.Printf("Transaction submitted: %s\n", tx.Hex())
	log.Println("Waiting for confirmation...")

	receipt, err := w.WaitForReceipt(ctx, tx)
	if err != nil {
		return fmt.Errorf("failed to get receipt: %w", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed (oracle may have reverted - check if time has passed and poll is registered)")
	}

	log.Printf("Poll closed successfully! Block: %d\n", receipt.BlockNumber.Uint64())
	return nil
}
