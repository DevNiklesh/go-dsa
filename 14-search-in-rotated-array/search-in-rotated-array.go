// https://takeuforward.org/data-structure/search-element-in-a-rotated-sorted-array/

package main

import "fmt"

// Brute force
// func searchInRotatedArray(nums []int, k int) int {
// 	for i, n := range nums {
// 		if n == k {
// 			return i
// 		}
// 	}
// 	return -1
// }

func searchInRotatedArray1(nums []int, k int) int {
	mid := 0
	low, high := 0, len(nums)-1

	for low <= high {
		mid = (low + high) / 2

		if nums[mid] == k {
			return mid
		}

		if nums[low] <= nums[mid] { // Left is sorted
			if nums[low] <= k && k < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] {
			if nums[mid] < k && k <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func main() {
	// fmt.Println(searchInRotatedArray([]int{4, 5, 6, 7, 0, 1, 2, 3}, 0))
	fmt.Println(searchInRotatedArray1([]int{4, 5, 6, 7, 0, 1, 2, 3}, 0))
	// fmt.Println(searchInRotatedArray1([]int{7, 8, 9, 1, 2, 3, 4, 5, 6}, 8))
}
