// https://takeuforward.org/data-structure/kadanes-algorithm-maximum-subarray-sum-in-an-array/

package main

import (
	"fmt"
	"math"
)

// Brute force
func maxSubArraySum(nums []int) int {
	maxSum := math.MinInt
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}

	}
	return maxSum
}

func maxSubArraySum1(nums []int) int {
	maxSum := 0
	start := 0
	ansStart, ansEnd := 0, 0

	sum := 0
	for i, n := range nums {
		if sum == 0 {
			start = i
		}

		sum += n
		if sum > maxSum {
			maxSum = sum

			ansStart = start
			ansEnd = i
		}

		if sum < 0 {
			sum = 0
		}
	}

	fmt.Println(nums[ansStart : ansEnd+1])
	return maxSum
}

func main() {
	fmt.Println(maxSubArraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArraySum([]int{1}))
	fmt.Println(maxSubArraySum([]int{}))
	fmt.Println(maxSubArraySum1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
