package main

import (
	"fmt"
	"sort"
)

func noOfMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	last := intervals[0]
	nonOverlaps := 1

	for _, curr := range intervals[1:] {
		if last[1] < curr[0] {
			last = curr
			nonOverlaps++
		}
	}

	return len(intervals) - nonOverlaps
}

func main() {
	fmt.Println(noOfMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(noOfMeetingRooms([][]int{{1, 18}, {18, 23}, {15, 29}, {4, 15}, {2, 11}, {5, 13}}))
}
