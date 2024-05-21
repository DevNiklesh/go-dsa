package main

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, new []int) [][]int {
	return merge(append(intervals, new))
}

func merge(intervals [][]int) [][]int {
	result := [][]int{}
	result = append(result, intervals[0])

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		top := result[len(result)-1]

		if interval[0] > top[1] {
			result = append(result, interval)
		} else {
			top[1] = max(interval[1], top[1])
		}
	}

	return result
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
