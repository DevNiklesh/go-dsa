// https://www.naukri.com/code360/problems/frog-jump_3621012?source=youtube&campaign=striver_dp_videos&utm_source=youtube&utm_medium=affiliate&utm_campaign=striver_dp_videos

package main

import (
	"fmt"
	"math"
)

func energy(stairs []int, n int, dp []int) int {
	if n == 0 {
		return 0
	}

	if dp[n] != -1 {
		return dp[n]
	}

	jump1 := energy(stairs, n-1, dp) + abs(stairs[n]-stairs[n-1])
	jump2 := math.MaxInt
	if n > 1 {
		jump2 = energy(stairs, n-2, dp) + abs(stairs[n]-stairs[n-2])
	}

	dp[n] = min(jump1, jump2)
	return dp[n]
}

// Tabular
func minEnergy(heights []int) int {
	n := len(heights)
	dp := make([]int, n)

	dp[0] = 0

	for i := 1; i < n; i++ {
		fs := dp[i-1] + abs(heights[i]-heights[i-1])
		ss := math.MaxInt
		if i > 1 {
			ss = dp[i-2] + abs(heights[i]-heights[i-2])
		}

		dp[i] = min(fs, ss)
	}

	return dp[n-1]
}

func minEnergy1(heights []int) int {
	n := len(heights)

	prev, prev2 := 0, 0

	for i := 1; i < n; i++ {
		fs := prev + abs(heights[i]-heights[i-1])
		ss := math.MaxInt
		if i > 1 {
			ss = prev2 + abs(heights[i]-heights[i-2])
		}
		prev2 = prev
		prev = min(fs, ss)
	}

	return prev
}

func main() {
	// stairs := []int{10, 20, 30, 10}
	stairs := []int{30, 10, 60, 10, 60, 50}
	dp := make([]int, len(stairs))
	for i := range dp {
		dp[i] = -1
	}

	fmt.Println(energy(stairs, len(stairs)-1, dp))
	fmt.Println(minEnergy(stairs))
	fmt.Println(minEnergy1(stairs))
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Tabular
// T: O(n*k)
// S: O(n)
func minEnergyWithKjumps(heights []int, k int) int {
	n := len(heights)
	dp := make([]int, n)

	dp[0] = 0

	for i := 1; i < n; i++ {
		minJ := math.MaxInt
		for j := 1; j <= k; j++ {
			if i-j >= 0 {
				jump := dp[i-j] + abs(heights[i]-heights[i-j])
				minJ = min(minJ, jump)
			}
		}

		dp[i] = minJ
	}

	return dp[n-1]
}
