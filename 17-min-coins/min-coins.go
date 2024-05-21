package main

import (
	"fmt"
	"math"
)

func f(i int, T int, arr []int, dp [][]int) int {
	if i == 0 {
		if T%arr[0] == 0 {
			return T / arr[0]
		}
		return 1e9
	}

	if dp[i][T] != -1 {
		return dp[i][T]
	}

	notTake := 0 + f(i-1, T, arr, dp)
	take := math.MaxInt
	if T >= arr[i] {
		take = 1 + f(i, T-arr[i], arr, dp)
	}

	dp[i][T] = min(take, notTake)
	return dp[i][T]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCoins(coins []int, target int) int {
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return f(len(coins)-1, target, coins, dp)
}

func main() {
	fmt.Println(minCoins([]int{1, 2, 3}, 8))
}
