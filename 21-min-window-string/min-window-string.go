package main

import (
	"fmt"
	"math"
)

func minWindowStr(s, t string) string {
	freq := make(map[byte]int)
	for i := range t {
		freq[t[i]]++
	}

	sidx := -1
	minLen := math.MaxInt
	count := 0
	l := 0

	for r := 0; r < len(s); r++ {
		if freq[s[r]] > 0 {
			count++
		}
		freq[s[r]]--

		for count == len(t) {
			if r-l+1 <= minLen {
				minLen = r - l + 1
				sidx = l
			}
			freq[s[l]]++
			if freq[s[l]] > 0 {
				count--
			}
			l++
		}
	}

	if sidx == -1 {
		return ""
	}
	return s[sidx:(sidx + minLen)]
}

func main() {
	fmt.Println(minWindowStr("ADOBECODEBANC", "ABC"))
}
