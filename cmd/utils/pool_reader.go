package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

// PoolReserves represents the reserves in a Uniswap V2 pool
type PoolReserves struct {
	Reserve0  *big.Int
	Reserve1  *big.Int
	Timestamp uint32
}

// GetPoolReserve reads the current reserves from a Uniswap V2 pool
func GetPoolReserve(client *ethclient.Client, poolAddress string) (*PoolReserves, error) {
	// For the demo, we'll return simulated values
	// In a real implementation, this would make a contract call

	// Simulated reserves (these would come from the blockchain in reality)
	reserves := &PoolReserves{
		Reserve0:  big.NewInt(10000000000000000), // 100 tokens (with 18 decimals)
		Reserve1:  big.NewInt(9500000000000000),  // 95 tokens (with 18 decimals)
		Timestamp: 12345678,
	}

	// In a full implementation, you would:
	// 1. Create a contract binding for the Uniswap V2 Pair interface
	// 2. Call the getReserves() function on the contract
	// 3. Parse and return the results

	return reserves, nil
}

// CalculateCurrentRatio calculates the current ratio of token0 and token1
func CalculateCurrentRatio(reserves *PoolReserves) float64 {

	// Convert big ints to float64 for ratio calculation
	reserve0Float := new(big.Float).SetInt(reserves.Reserve0)
	reserve1Float := new(big.Float).SetInt(reserves.Reserve1)

	// Calculate ratio
	ratio := new(big.Float).Quo(reserve0Float, reserve1Float)

	// Convert to float64
	result, _ := ratio.Float64()
	return result
}
