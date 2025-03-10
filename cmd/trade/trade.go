package trade

import (
	
	"github.com/spf13/cobra"
)

// Trade command

var TradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "Execute trades on Uniswap V2",
	Long: `The trade command allows you to execute trades on Uniswap V2 liquidity pools.
It supports swapping tokens to deliberately create imbalances in pools or regular trading.`,
}

func init() {
	// Persistant flags for all trade subcommands
	TradeCmd.PersistentFlags().StringP("rpc-url", "r", "https://eth-goerli.g.alchemy.com/v2/demo", "Ethereum RPC URL")
	TradeCmd.PersistentFlags().StringP("wallet", "w", "", "Wallet address to use for trades")
	TradeCmd.PersistentFlags().StringP("keystore-file", "k", "", "Path to keystore file")
	TradeCmd.PersistentFlags().String("gas-price", "auto", "Gas price in Gwei or 'auto'")
	TradeCmd.PersistentFlags().Uint64("gas-limit", 250000, "Gas limit for transactions")
}
