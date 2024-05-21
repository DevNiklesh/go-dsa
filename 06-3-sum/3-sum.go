package main

import (
	"fmt"
	"sort"
)

func findThreeSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1

		for j < k {
			s := nums[i] + nums[j] + nums[k]

			if s < target {
				j++
			} else if s > target {
				k--
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}

		}
	}

	return result
}

func main() {
	fmt.Println(findThreeSum([]int{-1, 0, 1, 2, -1, -4}, 0))
	fmt.Println(findThreeSum([]int{-1, 0, 1, 0}, 0))
}
