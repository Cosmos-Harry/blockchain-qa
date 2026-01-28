package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	privateKey string
	rpcURL     string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "poll-cli",
	Short: "CLI for interacting with the confidential poll dApp",
	Long: `A command-line interface for creating polls, voting with zero-knowledge proofs,
and viewing results in a privacy-preserving voting system.`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&privateKey, "private-key", "", "Ethereum private key (without 0x)")
	rootCmd.PersistentFlags().StringVar(&rpcURL, "rpc-url", "http://localhost:8545", "Ethereum RPC URL")

	// Mark private key as required for commands that need it
	// Commenting this out - it may cause issues
	// rootCmd.MarkPersistentFlagRequired("private-key")
}

func mustMarkRequired(cmd *cobra.Command, flag string) {
	if err := cmd.MarkFlagRequired(flag); err != nil {
		panic(err)
	}
}
