package main

import (
	"fmt"
)

func longestRepeatingChar(s string, k int) int {
	freq, maxFreq := make(map[string]int), 0
	maxLen := 0
	l := 0

	for r := 0; r < len(s); r++ {
		str := string(s[r])
		freq[str]++
		maxFreq = max(maxFreq, freq[str])

		if (r-l+1)-maxFreq > k {
			freq[string(s[l])]--
			l++
		}

		if (r-l+1)-maxFreq <= k {
			maxLen = max(maxLen, (r - l + 1))
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestRepeatingChar("AAABBCCDD", 2))
}
