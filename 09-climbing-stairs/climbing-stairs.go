// https://takeuforward.org/data-structure/dynamic-programming-climbing-stairs/

package main

import "fmt"

// recurssion
// T: O(n), space: O(n)+O(n) (array + recursion stack)
func climb(n int, dp []int) int {
	if n == 0 { // only one way to go to 0 steps (stand in 0th stair is one way)
		return 1
	}
	if n == 1 { // only one way to go to 1 step (from 0 only one way is there to reach 1st stair)
		return 1
	}

	if dp[n] != -1 {
		return dp[n]
	}

	dp[n] = climb(n-1, dp) + climb(n-2, dp)
	return dp[n]
}

// Forloop
// T: O(n) S: O(n)
func climb1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// forloop
// t: O(n) s: O(n)
func climb2(n int) int {
	prev2 := 1
	prev := 1

	for i := 2; i <= n; i++ {
		curi := prev + prev2
		prev2 = prev
		prev = curi
	}

	return prev
}

func main() {
	n := 6
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	fmt.Println(climb(n, dp))
	fmt.Println(climb1(n))
	fmt.Println(climb2(n))
}
