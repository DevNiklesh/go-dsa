package main

import (
	"fmt"
	"sort"
)

func canPersonAttend(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	last := intervals[0]

	for _, curr := range intervals[1:] {
		if curr[0] < last[1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canPersonAttend([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(canPersonAttend([][]int{{0, 8}, {8, 20}}))
	fmt.Println(canPersonAttend([][]int{{1, 18}, {18, 23}, {15, 29}, {4, 15}, {2, 11}, {5, 13}}))

}
