// https://takeuforward.org/data-structure/coin-change-2-dp-22/
package main

import "fmt"

func f(i int, T int, arr []int, dp [][]int) int {
	if i == 0 {
		if T%arr[i] == 0 {
			return 1
		}
		return 0
	}

	if dp[i][T] != -1 {
		return dp[i][T]
	}

	notTake := f(i-1, T, arr, dp)
	take := 0
	if T >= arr[i] {
		take += f(i, T-arr[i], arr, dp)
	}

	dp[i][T] = take + notTake
	return dp[i][T]
}

func noOfWaysR(coins []int, target int) int {
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return f(len(coins)-1, target, coins, dp)
}

func noOfWaysT(coins []int, target int) int {
	prev := make([]int, target+1)

	for i := 0; i < target+1; i++ {
		if i%coins[0] == 0 {
			prev[i] = 1
		}
	}

	for i := 1; i < len(coins); i++ {
		curr := make([]int, target+1)
		for T := 0; T < target+1; T++ {
			notTake := prev[T]
			take := 0
			if T >= coins[i] {
				take = curr[T-coins[i]]
			}
			curr[T] = take + notTake
		}
		prev = curr
	}

	return prev[target]
}

func main() {
	fmt.Println(noOfWaysR([]int{1, 2, 3}, 4))
	fmt.Println(noOfWaysT([]int{1, 2, 3}, 4))
}
