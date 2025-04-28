package arbitrage

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	autoInterval  uint
	maxExecutions int
	autoTimeLimit uint
	minProfitAuto float64
)

// The auto command for arbitrage
var AutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Automatically scan and execute arbitrage trades",
	Long: `Run the arbitrage bot in automatic mode, continuously scanning for opportunities and executing trades when profitable opportunities are found. Set minimum profit thresholds and other safety parameters to control execution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get persistent flags
		rpcURL, _ := cmd.Flags().GetString("rpc-url")
		wallet, _ := cmd.Flags().GetString("wallet")
		keystoreFile, _ := cmd.Flags().GetString("keystore-file")
		gasPrice, _ := cmd.Flags().GetString("gas-price")
		gasLimit, _ := cmd.Flags().GetUint64("gas-limit")

		// Override min-profit if specified specifically for auto mode
		minProfit, _ := cmd.Flags().GetFloat64("min-profit")
		if cmd.Flags().Changed("auto-min-profit") {
			minProfit = minProfitAuto
		}

		fmt.Println("🤖 Starting arbitrage bot in AUTO mode...")
		fmt.Printf("  RPC URL: %s\n", rpcURL)
		fmt.Printf("  Wallet: %s\n", wallet)
		fmt.Printf("  Keystore: %s\n", keystoreFile)
		fmt.Printf("  Scan Interval: %d seconds\n", autoInterval)
		fmt.Printf("  Min Profit: %.2f%%\n", minProfit)
		fmt.Printf("  Gas Price: %s\n", gasPrice)
		fmt.Printf("  Gas Limit: %d\n", gasLimit)

		if maxExecutions > 0 {
			fmt.Printf("  Max Executions: %d\n", maxExecutions)
		} else {
			fmt.Println("  Max Executions: Unlimited")
		}

		if autoTimeLimit > 0 {
			fmt.Printf("  Time Limit: %d minutes\n", autoTimeLimit)
		} else {
			fmt.Println("  Time Limit: None (running until stopped)")
		}

		fmt.Println("\n⚠️ Press Ctrl+C to stop the bot")
		fmt.Println("\n🔄 Bot started at", time.Now().Format(time.RFC3339))

		// Set up signal handling for graceful shutdown
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		// Create a ticker for the scan inerval
		ticker := time.NewTicker(time.Duration(autoInterval) * time.Second)
		defer ticker.Stop()

		// Create a timeout if specified
		var timeout <-chan time.Time
		if autoTimeLimit > 0 {
			timeout = time.After(time.Duration(autoTimeLimit) * time.Minute)
		}

		// Start the scanning and execution loop
		executionCount := 0
		scanCount := 0

		for {
			select {
			case <-ticker.C:
				scanCount++
				fmt.Printf("\n[%s] Scan #%d: Checking for arbitrage opportunities...\n",
					time.Now().Format("15:04:05"), scanCount)

				// Simulate finding an opportunity (in reality, this would check pool states)
				profitFound := scanCount%4 == 0 // Just a quick demo

				if profitFound {
					profit := minProfit + float64(scanCount%5)
					fmt.Printf("✅ Opportunity found! Potential profit: %.2f%%\n", profit)

					if profit >= minProfit {
						executionCount++
						fmt.Printf("💰 Executing arbitrage trade #%d\n", executionCount)
						time.Sleep(2 * time.Second)                                         // Simulate execution time
						fmt.Printf("✅ Trade executed! Actual profit: %.2f%%\n", profit*0.9) // Slightly less due to slippage

						if maxExecutions > 0 && executionCount >= maxExecutions {
							fmt.Printf("\n🛑 Reached maximum number of executions (%d)\n", maxExecutions)
							return
						}
					} else {
						fmt.Printf("⚠️ Profit too low (%.2f%% < %.2f%%). Skipping execution.\n",
							profit, minProfit)
					}
				} else {
					fmt.Println("No profitable opportunities found in this scan")
				}

			case <-timeout:
				if autoTimeLimit > 0 {
					fmt.Printf("\n⏱️ Auto mode time limit (%d minutes) reached\n", autoTimeLimit)
					fmt.Printf("Summary: %d scans, %d executions\n", scanCount, executionCount)
					return
				}

			case <-sigs:
				fmt.Println("\n\n🛑 Received termination signal. Shutting down...")
				fmt.Printf("Summary: %d scans, %d executions\n", scanCount, executionCount)
				return
			}
		}
	},
}

func init() {
	// Add command-specific flags
	AutoCmd.Flags().UintVar(&autoInterval, "interval", 30, "Scan interval in seconds")
	AutoCmd.Flags().IntVar(&maxExecutions, "max-executions", 0, "Maximum number of trades to execute (0 for unlimited)")
	AutoCmd.Flags().UintVar(&autoTimeLimit, "time-limit", 0, "Time limit in minutes (0 for no limit)")
	AutoCmd.Flags().Float64Var(&minProfitAuto, "auto-min-profit", 0, "Minimum profit percentage override for auto mode")

	// Add this command to the parent arbitrage command
	ArbitrageCmd.AddCommand(AutoCmd)
}
