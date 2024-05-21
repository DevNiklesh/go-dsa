package main

import (
	"fmt"
)

func repeatingMissingNo(nums []int) []int {
	a, b := -1, -1
	length := len(nums)
	freq := make([]int, length)

	for _, val := range nums {
		freq[val-1]++
	}

	for i, val := range freq {
		if val == 0 {
			b = i + 1
		}
		if val == 2 {
			a = i + 1
		}
	}

	if b == -1 {
		b = len(nums)
	}

	return []int{a, b}
}

func main() {
	fmt.Println(repeatingMissingNo([]int{3, 1, 2, 5, 4, 6, 7, 5}))
}
