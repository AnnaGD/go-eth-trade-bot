package trade

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	tokenIn     string
	tokenOut    string
	amount      string
	slippage    float64
	deadlineMin uint
)

// Execute command for trading
var ExecuteCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute a trade on Uniswap V2",
	Long: `Execute a token swap on Uniswap V2 liquidity pools.
This command allows you to swap one token for another, specifying
amount, slippage tolerance, and deadline.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Persistant falgs
		rpcURL, _ := cmd.Flags().GetString("rpc-url")
		wallet, _ := cmd.Flags().GetString("wallet")
		keystoreFile, _ := cmd.Flags().GetString("keystore-file")
		gasPrice, _ := cmd.Flags().GetString("gas-price")
		gasLimit, _ := cmd.Flags().GetUint64("gas-limit")

		fmt.Println("🔄 Execute trade ...")
		fmt.Printf("  Token In: %s\n", tokenIn)
		fmt.Printf("  Token Out: %s\n", tokenOut)
		fmt.Printf("  Amount: %s\n", amount)
		fmt.Printf("  Slippage: %.2f%%\n", slippage)
		fmt.Printf("  Deadline: %d minutes\n", deadlineMin)
		fmt.Printf("  RPC URL: %s\n", rpcURL)
		fmt.Printf("  Wallet: %s\n", wallet)
		fmt.Printf("  Keystore: %s\n", keystoreFile)
		fmt.Printf("  Gas Price: %s\n", gasPrice)
		fmt.Printf("  Gas Limit: %d\n", gasLimit)

		// TODO: Implement actual trade execution logic
		fmt.Println("⚠ Trade execusion not yet implemented")
	},
}

func init() {
	// Command specific flags
	ExecuteCmd.Flags().StringVar(&tokenIn, "token-in", "ETH", "Input token symbol or address")
	ExecuteCmd.Flags().StringVar(&tokenOut, "token-out", "", "Output token symbol or address")
	ExecuteCmd.Flags().StringVar(&amount, "amount", "", "Amount of input token to swap")
	ExecuteCmd.Flags().Float64Var(&slippage, "slippage", 0.5, "Slippage tolerance percentage")
	ExecuteCmd.Flags().UintVar(&deadlineMin, "deadline", 20, "Transaction deadline in minutes")

	// Required flags
	ExecuteCmd.MarkFlagRequired()
	ExecuteCmd.MarkFlagRequired()

	// Additional command to the parent trade command
	TradeCmd.AddCommand(ExecuteCmd)
}
