/*
	This file represents the core scanning functionality of the arbitrage bot, providing a mechanism to continuously monitor pools and identify profitable trading opportunities based on price imbalances. While currently using simulated data, it's structured to be easily upgraded to interact with real blockchain data.

*/


package arbitrage

import (
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AnnaGD/go-eth-trade-bot/cmd/constants" 
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var (
	scanInterval  uint
	outputFormat  string
	scanTimeLimit uint
	selectedPools []string
)

// PoolReserves represents the reserves in a Uniswap V2 pool
type PoolReserves struct {
	Reserve0  *big.Int
	Reserve1  *big.Int
	Timestamp uint32
}

// Defines the main scan command with its usage, descriptions, and run function.
/*
	Retrieves configuration from command flags, Establishes an Ethereum client connection,
	Sets up pool monitoring with specified intervals, Implements graceful termination handling
	Runs the main scanning loop that:
		Checks each pool's reserves
		Calculates current token ratios
		Compares to target ratios
		Identifies and reports arbitrage opportunities
		Calculates potential profit percentages
*/
var ScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for arbitrage opportunities",
	Long: `Scan Uniswap V2 pools for potential arbitrage opportunities. This command monitors pool states and identifies imbalances that could be exploited for profit.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get persistent flags
		rpcURL, _ := cmd.Flags().GetString("rpc-url")
		minProfit, _ := cmd.Flags().GetFloat64("min-profit")

		fmt.Println("🔍 Starting to scan pools for arbitrage opportunities...")
		fmt.Printf("  RPC URL: %s\n", rpcURL)
		fmt.Printf("  Min Profit: %.2f%%\n", minProfit)
		fmt.Printf("  Scan Interval: %d seconds\n", scanInterval)

		if scanTimeLimit > 0 {
			fmt.Printf("  Time Limit: %d minutes\n", scanTimeLimit)
		} else {
			fmt.Println("  Time Limit: None (running until stopped)")
		}

		// Connect to Ethereum
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			fmt.Printf("Error connecting to Ethereum: %v\n", err)
			return
		}

		// If no pools selected, use all pools
		if len(selectedPools) == 0 {
			for pool := range constants.UniV2Pools {
				selectedPools = append(selectedPools, pool)
			}
		}
		
		fmt.Printf("Monitoring %d pools with %d second interval\n", 
			len(selectedPools), scanInterval)

		// Setup for graceful termination
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		// Start timing
		startTime := time.Now()
		fmt.Println("⏳ Scanning started at", startTime.Format(time.RFC3339))
		fmt.Println("Press Ctrl+C to stop scanning")

		// Create a timer for the scan interval
		ticker := time.NewTicker(time.Duration(scanInterval) * time.Second)
		defer ticker.Stop()

		// Create a timeout if specified
		var timeout <-chan time.Time
		if scanTimeLimit > 0 {
			timeout = time.After(time.Duration(scanTimeLimit) * time.Minute)
		}

		// Run scan loop
		opportunityCount := 0
		scanCount := 0

		for {
			select {
			case <-ticker.C:
				scanCount++
				fmt.Printf("\n[%s] Scan #%d: Checking for arbitrage opportunities...\n",
					time.Now().Format("15:04:05"), scanCount)

				// Check each selected pool
				for _, poolName := range selectedPools {
					poolAddress, exists := constants.UniV2Pools[poolName]
					if !exists {
						fmt.Printf("Pool %s not found\n", poolName)
						continue
					}
					
					targetRatio, hasTarget := constants.TargetRatios[poolName]
					if !hasTarget {
						fmt.Printf("No target ratio for %s, skipping\n", poolName)
						continue
					}
					
					// Get pool reserves
					reserves, err := getPoolReserves(client, poolAddress)
					if err != nil {
						fmt.Printf("Error reading %s: %v\n", poolName, err)
						continue
					}
					
					// Calculate current ratio
					currentRatio := calculateCurrentRatio(reserves)
					
					// Calculate imbalance percentage
					imbalancePercent := ((currentRatio - targetRatio) / targetRatio) * 100
					
					fmt.Printf("%s: Current Ratio: %.4f (Target: %.4f)\n", 
						poolName, currentRatio, targetRatio)
					
					// Check if there's a significant imbalance
					if abs(imbalancePercent) > 1.0 {  // More than 1% imbalance
						fmt.Printf("  ✅ OPPORTUNITY: %.2f%% imbalance detected!\n", imbalancePercent)
						
						// Calculate potential profit
						profitPercent := calculatePotentialProfit(currentRatio, targetRatio)
						
						// Check if profit meets minimum threshold
						if profitPercent >= minProfit {
							opportunityCount++
							fmt.Printf("  💰 Opportunity #%d - Potential Profit: %.2f%%\n", 
								opportunityCount, profitPercent)
							
							// In a full implementation, you would:
							// 1. Calculate optimal trade amounts
							// 2. Execute the trade if in auto mode
							// 3. Log the opportunity details
						} else {
							fmt.Printf("  ❌ Profit too low: %.2f%% (min: %.2f%%)\n", profitPercent, minProfit)
						}
					} else {
						fmt.Printf("  ❌ No significant imbalance\n")
					}
				}
				
			case <-timeout:
				if scanTimeLimit > 0 {
					fmt.Printf("\n\n⏱️ Scan time limit (%d minutes) reached\n", scanTimeLimit)
					fmt.Printf("Found %d opportunities in %d scans\n", opportunityCount, scanCount)
					return
				}

			case <-sigs:
				fmt.Println("\n\n🛑 Received termination signal. Shutting down...")
				fmt.Printf("Found %d opportunities in %d scans\n", opportunityCount, scanCount)
				return
			}
		}
	},
}

// This is a placeholder function that simulates getting pool reserves

/*
	Currently returns deterministic mock values based on pool addresses
	Is meant to be replaced with actual Uniswap V2 contract calls in production
	Returns a `PoolReserves` struct with reserve values and timestamp
*/

func getPoolReserves(client *ethclient.Client, poolAddress string) (*PoolReserves, error) {
	
	// Convert address string to common.Address
	addr := common.HexToAddress(poolAddress)
	
	// Create deterministic values based on the address
	// This creates different values for different pools but stays consistent
	reserve0 := new(big.Int).SetBytes(addr.Bytes()[:16])
	reserve0.Mod(reserve0, big.NewInt(1000))
	reserve0.Add(reserve0, big.NewInt(500))
	reserve0.Mul(reserve0, big.NewInt(1000000000000000000)) // 18 decimals
	
	reserve1 := new(big.Int).SetBytes(addr.Bytes()[16:])
	reserve1.Mod(reserve1, big.NewInt(1000))
	reserve1.Add(reserve1, big.NewInt(500))
	reserve1.Mul(reserve1, big.NewInt(1000000000000000000)) // 18 decimals
	
	reserves := &PoolReserves{
		Reserve0:  reserve0,
		Reserve1:  reserve1,
		Timestamp: uint32(time.Now().Unix()),
	}
	
	return reserves, nil
}

// calculateCurrentRatio calculates the current ratio of token0 to token1
func calculateCurrentRatio(reserves *PoolReserves) float64 {

	// Convert big ints to float64 for ratio calculation
	reserve0Float := new(big.Float).SetInt(reserves.Reserve0)
	reserve1Float := new(big.Float).SetInt(reserves.Reserve1)
	
	// Calculate ratio
	ratio := new(big.Float).Quo(reserve0Float, reserve1Float)
	
	// Convert to float64
	result, _ := ratio.Float64()
	return result
}

// abs returns the absolute value of x
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// calculatePotentialProfit calculates potential profit from an arbitrage
func calculatePotentialProfit(currentRatio, targetRatio float64) float64 {
	// Profit model: larger imbalance = higher profit potential
	imbalancePercent := abs((currentRatio - targetRatio) / targetRatio) * 100
	
	// Discount for transaction costs
	return imbalancePercent * 0.7  // 70% of imbalance becomes profit (after gas)
}

func init() {
	// Add command-specific flags

	// Scan frequency in seconds (default 30)
	ScanCmd.Flags().UintVar(&scanInterval, "interval", 30, "Scan interval in seconds")

	// Specific pools to monitor (default: all)
	ScanCmd.Flags().StringSliceVar(&selectedPools, "pools", []string{}, "Pool names to monitor (default: all)")

	// Output format (text or JSON)
	ScanCmd.Flags().StringVar(&outputFormat, "output", "text", "Output format (text, json)")

	// Duration to run the scan (0 for unlimited)
	ScanCmd.Flags().UintVar(&scanTimeLimit, "time-limit", 0, "Time limit for scanning in minutes (0 for no limit)")

	// Add this command to the parent arbitrage command
	ArbitrageCmd.AddCommand(ScanCmd)
}