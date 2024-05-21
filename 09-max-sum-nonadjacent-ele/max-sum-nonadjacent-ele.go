package main

import "fmt"

// Recurssion
func maxSumAdjacentEle(i int, arr []int, dp []int) int {
	if i < 0 {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}

	pick := arr[i] + maxSumAdjacentEle(i-2, arr, dp)
	notPick := 0 + maxSumAdjacentEle(i-1, arr, dp)

	dp[i] = max(pick, notPick)
	return dp[i]
}

// Tabular
func maxSumAdjacentEle1(arr []int) int {
	dp := make([]int, len(arr))

	dp[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		take := arr[i]
		if i > 1 {
			take += dp[i-2]
		}
		notTake := 0 + dp[i-1]

		dp[i] = max(take, notTake)
	}

	return dp[len(arr)-1]
}

// Tabular space optimized
func maxSumAdjacentEle2(arr []int) int {
	prev2 := 0
	prev := arr[0]

	for i := 1; i < len(arr); i++ {
		take := arr[i]
		if i > 1 {
			take += prev2
		}
		notTake := prev

		curri := max(take, notTake)
		prev2 = prev
		prev = curri
	}

	return prev
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := []int{2, 1, 4, 9}
	dp := make([]int, len(n))
	for i := range dp {
		dp[i] = -1
	}
	fmt.Println(maxSumAdjacentEle(len(n)-1, n, dp))
	fmt.Println(maxSumAdjacentEle1(n))
	fmt.Println(maxSumAdjacentEle2(n))
}
