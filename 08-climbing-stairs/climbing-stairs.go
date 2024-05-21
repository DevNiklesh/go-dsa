package main

import "fmt"

// Recurssion
func maxWaysToClimbStairsRecurssion(n int, dp []int) int {
	fmt.Println(dp)
	if n <= 1 {
		return 1
	}
	if dp[n] != -1 {
		return dp[n]
	}

	dp[n] = maxWaysToClimbStairsRecurssion(n-1, dp) + maxWaysToClimbStairsRecurssion(n-2, dp)
	return dp[n]
}

func maxWaysToClimbStairsTabular(n int) int {
	prev2 := 1
	prev := 1

	for i := 2; i < n+1; i++ {
		curri := prev2 + prev
		prev2 = prev
		prev = curri
	}

	return prev
}

func main() {
	// Recurrsion
	n := 4
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	fmt.Println(maxWaysToClimbStairsRecurssion(n, dp))

	// fmt.Println(maxWaysToClimbStairsTabular(2))
	// fmt.Println(maxWaysToClimbStairsTabular(3))
	fmt.Println(maxWaysToClimbStairsTabular(4))
}
