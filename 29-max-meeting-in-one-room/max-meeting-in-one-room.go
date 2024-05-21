package main

import "fmt"

type interval struct {
	id       int
	schedule []int
}

func maxMeetings(s, f []int) []int {
	intervals := make([]interval, len(f))
	for i := range s {
		intervals[i] = interval{i + 1, []int{s[i], f[i]}}
	}

	last := intervals[0]
	result := []int{last.id}

	for _, currInterval := range intervals {
		if currInterval.schedule[0] >= last.schedule[1] {
			last = currInterval
			result = append(result, currInterval.id)
		}
	}

	return result
}

func main() {
	fmt.Println(maxMeetings([]int{1, 3, 0, 5, 8, 5}, []int{2, 4, 6, 7, 9, 9}))
}
