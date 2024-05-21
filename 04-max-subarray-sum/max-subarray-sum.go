package main

import "fmt"

func maxSubArraySum(nums []int) (int, []int) {
	maxSum := 0

	sum := 0
	start := 0
	ansStart, ansEnd := 0, 0
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

	return maxSum, []int{ansStart, ansEnd}
}
func main() {
	fmt.Println(maxSubArraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
