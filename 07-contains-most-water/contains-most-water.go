package main

import (
	"fmt"
	"math"
)

func mostWater(nums []int) int {
	l, r := 0, len(nums)-1

	maxArea := math.MinInt

	for l < r {
		area := (r - l) * min(nums[l], nums[r])
		maxArea = max(maxArea, area)

		if l < r {
			l++
		} else {
			r--
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(mostWater([]int{1, 5, 4, 3}))
	fmt.Println(mostWater([]int{3, 1, 2, 4, 5}))
}
