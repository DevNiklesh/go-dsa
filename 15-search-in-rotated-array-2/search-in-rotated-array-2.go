package main

import "fmt"

func searchInRotatedArray(nums []int, k int) int {
	mid := 0
	low, high := 0, len(nums)-1

	for low <= high {
		mid = (low + high) / 2

		if nums[mid] == k {
			return mid
		}

		if nums[low] == nums[mid] && nums[mid] == nums[high] {
			low += 1
			high -= 1
			continue
		}

		if nums[low] <= nums[mid] { // left is sorted
			if nums[low] <= k && k < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
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
	fmt.Println(searchInRotatedArray([]int{7, 8, 1, 2, 3, 3, 3, 4, 5, 6}, 3))
	fmt.Println(searchInRotatedArray([]int{3, 1, 2, 3, 3, 3, 3}, 2))
	fmt.Println(searchInRotatedArray([]int{7, 8, 1, 2, 3, 3, 3, 4, 5, 6}, 10))
}
