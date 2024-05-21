package main

import (
	"fmt"
	"math"
)

func f(i int, T int, arr []int, dp [][]int) int {
	if i == 0 {
		if T%arr[i] == 0 {
			return T / arr[i]
		}
		return math.MaxInt
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

func minCoinsR(arr []int, T int) int {
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, T+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return f(len(arr)-1, T, arr, dp)
}

func minCoinsT(arr []int, T int) int {
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, T+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return f(len(arr)-1, T, arr, dp)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCoinsR([]int{1, 2, 3}, 8))
	fmt.Println(minCoinsT([]int{1, 2, 3}, 8))
}
