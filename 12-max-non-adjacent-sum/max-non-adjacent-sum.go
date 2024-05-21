// https://takeuforward.org/data-structure/maximum-sum-of-non-adjacent-elements-dp-5/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxNonAdjacentSum(idx int, nums []int, dp []int) int {
	if idx == 0 {
		return nums[idx]
	}
	if idx < 0 {
		return 0
	}

	if dp[idx] != -1 {
		return dp[idx]
	}

	pick := nums[idx] + maxNonAdjacentSum(idx-2, nums, dp)
	notPick := 0 + maxNonAdjacentSum(idx-1, nums, dp)

	dp[idx] = max(pick, notPick)
	return dp[idx]
}

func main() {
	nums := []int{2, 1, 4, 9}
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = -1
	}
	fmt.Println(maxNonAdjacentSum(len(nums)-1, nums, dp))
}
