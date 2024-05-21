package main

import (
	"fmt"
	"math"
	"sort"
)

// T: O(nlogn) | s: O(1)
func findMissingNumber(nums []int) int {
	sort.Ints(nums)

	missing := math.MinInt
	for i, n := range nums {
		fmt.Println(i, n)
		if i != n {
			return i
		}
	}

	if missing == math.MinInt {
		missing = len(nums)
	}

	return missing
}

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	arrSum := 0

	for _, num := range nums {
		arrSum += num
	}

	return sum - arrSum
}

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	// fmt.Println(findMissingNumber([]int{1, 2}))
}
