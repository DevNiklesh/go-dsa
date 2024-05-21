package main

import "fmt"

// Brute force
func maxProductSubarray(nums []int) int {
	n := len(nums)
	maxProduct := 1

	leftP, rightP := 1, 1

	for i := 0; i < len(nums); i++ {
		leftP *= nums[i]
		rightP *= nums[n-1-i]

		maxProduct = max(maxProduct, max(leftP, rightP))

		if leftP == 0 {
			leftP = 1
		}

		if rightP == 0 {
			rightP = 1
		}
	}

	return maxProduct
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProductSubarray([]int{1, 2, 3, 4, 5, 0}))    // 120
	fmt.Println(maxProductSubarray([]int{1, 2, -3, 0, -4, -5})) // 20
}
