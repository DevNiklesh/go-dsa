// https://takeuforward.org/data-structure/contains-duplicate-check-if-a-value-appears-atleast-twice/

package main

import (
	"fmt"
	"sort"
)

func isDuplicatePresent(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(isDuplicatePresent([]int{1, 2, 3, 4}))
}
