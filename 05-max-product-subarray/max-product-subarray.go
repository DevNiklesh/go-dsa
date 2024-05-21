package main

import (
	"fmt"
	"math"
)

func maxProductSubArray(nums []int) int {
	length := len(nums)
	maxProduct := math.MinInt

	pre, suffix := 1, 1
	for i, n := range nums {
		if pre == 0 {
			pre = 1
		}
		if suffix == 0 {
			suffix = 1
		}

		pre *= n
		suffix *= nums[length-1-i]
		if pre > suffix {
			if maxProduct < pre {
				maxProduct = pre
			}
		} else {
			if maxProduct < suffix {
				maxProduct = suffix
			}
		}
	}

	return maxProduct
}

func main() {
	// fmt.Println(maxProductSubArray([]int{1, 2, -3, 0, -4, -5}))
	// fmt.Println(maxProductSubArray([]int{-2, 0, -1}))
	fmt.Println(maxProductSubArray([]int{2, 3, -2, 4}))
	// fmt.Println(maxProductSubArray([]int{-2, 3, 4, -1, 0, -2, 3, 1, 4, 0, 4, 6, -1, 4}))

}
