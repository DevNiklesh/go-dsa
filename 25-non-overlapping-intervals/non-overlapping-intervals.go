package main

import (
	"fmt"
	"sort"
)

func eraceOverlappingIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)

	nonOverlaps := 1
	last := intervals[0]

	for _, interval := range intervals[1:] {
		if last[1] <= interval[0] {
			last = interval
			nonOverlaps++
		}
	}

	return len(intervals) - nonOverlaps
}

func main() {
	fmt.Println(eraceOverlappingIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(eraceOverlappingIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	fmt.Println(eraceOverlappingIntervals([][]int{{1, 2}, {2, 3}}))
}
