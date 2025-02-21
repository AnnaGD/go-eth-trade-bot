package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
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
