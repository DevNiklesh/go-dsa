package main

import (
	"fmt"
)

func longestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxLen, l := 0, 0

	seen := make(map[string]int)
	for r := 0; r < len(s); r++ {
		str := string(s[r])
		// fmt.Println("L:", l, "R:", r, "maxLen:", maxLen, "Seen:", seen)
		if isFound(seen, str) { // seen
			for l < r && isFound(seen, str) {
				delete(seen, string(s[l]))
				l++
			}
		}

		seen[str] = r
		maxLen = max(maxLen, r-l+1)
	}

	return maxLen
}

func isFound(hash map[string]int, s string) bool {
	if _, ok := hash[s]; ok {
		return ok
	}
	return false
}

func main() {
	fmt.Println(longestSubstring("abcabcbb"))
	fmt.Println(longestSubstring("bbbbbbb"))
	fmt.Println(longestSubstring("pwwkew"))
}
