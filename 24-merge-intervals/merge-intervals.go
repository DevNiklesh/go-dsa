package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	result := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result = append(result, intervals[0])
	for _, interval := range intervals {
		last := result[len(result)-1]

		if interval[0] > last[1] {
			result = append(result, interval)
		} else if interval[1] > last[1] {
			last[1] = interval[1]
		}
	}

	return result
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
