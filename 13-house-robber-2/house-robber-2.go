// https://takeuforward.org/data-structure/dynamic-programming-house-robber-dp-6/

package main

import "fmt"

func maxMoneyR(idx int, arr []int, dp []int) int {
	if idx == 0 {
		return arr[idx]
	}
	if idx < 0 {
		return 0
	}

	if dp[idx] != -1 {
		return dp[idx]
	}

	take := arr[idx] + maxMoneyR(idx-2, arr, dp)
	notTake := 0 + maxMoneyR(idx-1, arr, dp)

	dp[idx] = max(take, notTake)
	return dp[idx]
}

func maxMoneyT(arr []int) int {
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

func robStreet(arr []int) int {
	ans1 := maxMoneyT(arr[1:])
	ans2 := maxMoneyT(arr[:len(arr)-1])

	return max(ans1, ans2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{2, 1, 4, 9}
	// arr := []int{1, 2, 3, 1}
	fmt.Println(maxMoneyT(arr))

	dp := make([]int, len(arr))
	for i := range arr {
		dp[i] = -1
	}
	fmt.Println(maxMoneyR(len(arr)-1, arr, dp))
	fmt.Println(robStreet(arr))
}
