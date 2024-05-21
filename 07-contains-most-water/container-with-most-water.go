// https://www.geeksforgeeks.org/container-with-most-water/

package main

import "fmt"

// Brute force
func mostWater(nums []int) int {
	maxWater := 0

	distance, height := 1, 1
	for i, n := range nums {

		for j := i + 1; j < len(nums); j++ {
			distance = j - i
			if n < nums[j] {
				height = n
			} else {
				height = nums[j]
			}
			area := distance * height
			if area > maxWater {
				maxWater = area
			}
		}
	}

	return maxWater
}

func mostWater1(nums []int) int {
	maxArea := 0

	l, r := 0, len(nums)-1

	for l < r {
		area := min(nums[l], nums[r]) * (r - l)
		if area > maxArea {
			maxArea = area
		}

		if nums[l] > nums[r] {
			r--
		} else {
			l++
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(mostWater([]int{1, 5, 4, 3}))
	fmt.Println(mostWater1([]int{3, 1, 2, 4, 5}))
	fmt.Println(mostWater1([]int{8, 5, 9, 1, 10, 2, 6}))
}
