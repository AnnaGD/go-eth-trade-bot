package cmd

import (
	"fmt"
	"os"

	"github.com/AnnaGD/go-eth-trade-bot/cmd/keystore"
	"github.com/AnnaGD/go-eth-trade-bot/cmd/keystore/arbitrage"
	"github.com/AnnaGD/go-eth-trade-bot/cmd/trade"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tradebot",
	Short: "An automated arbitrage trading bot for EVM-compatible networks",
	Long: `Tradebot scans multiple decentralized exchanges (DEXs) across EVM-compatible testnets 
and mainnets, detecting arbitrage opportunities. It automates trade execution based on real-time 
price discrepancies, optimizing transaction profitability.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

	// Global persistent flags
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose output")

	// Top level commands
	rootCmd.AddCommand(keystore.KeystoreCmd)
	rootCmd.AddCommand(trade.TradeCmd)
	rootCmd.AddCommand(arbitrage.ArbitrageCmd)
}
