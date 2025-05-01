package cmd

import (
	"fmt"
	"os"

	"github.com/AnnaGD/go-eth-trade-bot/cmd/keystore"
	"github.com/AnnaGD/go-eth-trade-bot/cmd/trade/arbitrage"
	"github.com/AnnaGD/go-eth-trade-bot/cmd/trade"
	"github.com/spf13/cobra"
)

// Main function to execute the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Root command for the CLI
var rootCmd = &cobra.Command{
	Use:   "tradebot",
	Short: "An automated arbitrage trading bot for EVM-compatible networks",
	Long: `Tradebot scans multiple decentralized exchanges (DEXs) across EVM-compatible testnets and mainnets, detecting arbitrage opportunities. It automates trade execution based on real-time 
price discrepancies, optimizing transaction profitability.`,
}

// Initialize the `rootCmd` with addtional top level subcommands
func init() {

	// Global persistent flags
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose output")

	// Keystore management for secret keys
	rootCmd.AddCommand(keystore.KeystoreCmd)

	// Trade command
	rootCmd.AddCommand(trade.TradeCmd)

	// Arbitrage command
	rootCmd.AddCommand(arbitrage.ArbitrageCmd)
}
