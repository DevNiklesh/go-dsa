package main

import (
	"fmt"
	"math"
)

func minInRotatedArr(nums []int) int {
	mid, l, h := 0, 0, len(nums)-1

	m := math.MaxInt
	for l <= h {
		mid = (l + h) / 2

		if nums[l] <= nums[h] {
			m = min(m, nums[l])
			break
		}

		if nums[l] <= nums[mid] { // left is sorted
			m = min(m, nums[l])
			l = mid + 1
		} else {
			m = min(m, nums[mid])
			h = mid - 1
		}
	}

	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minInRotatedArr([]int{4, 5, 6, 7, 0, 1, 2, 3}))
	fmt.Println(minInRotatedArr([]int{3, 4, 5, 1, 2}))
}
