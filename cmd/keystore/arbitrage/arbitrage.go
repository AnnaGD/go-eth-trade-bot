package arbitrage

import (
	"fmt"
    "github.com/AnnaGD/go-eth-trade-bot/cmd/constants"
    "github.com/spf13/cobra"
)

// Arbitrage command
var ArbitrageCmd = &cobra.Command{
	Use:   "arbitrage",
	Short: "Detect and execute arbitrage opportunities",
	Long: `The arbitrage command allows you to scan for and execute
profitable arbitrage opportunities between different Uniswap V2 pools.
It can detect imbalances and automatically execute trades to capitalize on price differences.`,
}

func init() {

	ArbitrageCmd.AddCommand(ScanCmd)
	ArbitrageCmd.AddCommand(ExecuteCmd)
	ArbitrageCmd.AddCommand(AutoCmd)


	// Persistent flags for all arbitrage subcommands
	ArbitrageCmd.PersistentFlags().StringP("rpc-url", "r", "https://eth-goerli.g.alchemy.com/v2/demo", "Ethereum RPC URL")
	ArbitrageCmd.PersistentFlags().StringP("wallet", "w", "", "Wallet address to use for arbitrage")
	ArbitrageCmd.PersistentFlags().StringP("keystore-file", "k", "", "Path to keystore file")
	ArbitrageCmd.PersistentFlags().Float64P("min-profit", "p", 0.5, "Minimum profit percentage")
	ArbitrageCmd.PersistentFlags().String("gas-price", "auto", "Gas price in Gwei or 'auto'")
	ArbitrageCmd.PersistentFlags().Uint64("gas-limit", 350000, "Gas limit for transactions")

	// Display available pools
	fmt.Println("Available pools for arbitrage: ")
	for poolName, address := range constants.UniV2Pools {
		targetRatio, exists := constants.TargetRatios[poolName]
		if exists {
			fmt.Printf("  %s (%s) - Target Ratio: %.2f\n", poolName, address[:10]+"...", targetRatio)
		} else {
			fmt.Printf("  %s (%s)\n", poolName, address[:10]+"...")
		}
	}

}
