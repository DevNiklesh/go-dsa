// https://takeuforward.org/data-structure/two-sum-check-if-a-pair-with-given-sum-exists-in-array/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, n := range nums {
		balance := target - n
		if _, ok := hashmap[balance]; ok {
			return []int{hashmap[balance], i}
		}
		hashmap[n] = i
	}
	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 6, 5, 8, 11}, 14))
}
