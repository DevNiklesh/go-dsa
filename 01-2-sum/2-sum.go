package main

import "fmt"

func findTowsum(arr []int, target int) []int {
	result := []int{}

	hash := make(map[int]int)
	for i, n := range arr {
		if _, ok := hash[target-n]; ok {
			return []int{hash[target-n], i}
		}
		hash[n] = i
	}

	return result
}

func main() {
	fmt.Println(findTowsum([]int{2, 6, 5, 8, 11}, 14))
}
