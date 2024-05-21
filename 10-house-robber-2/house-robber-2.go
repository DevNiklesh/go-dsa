package main

import "fmt"

func maxSteel(arr []int) int {
	prev := arr[0]
	prev2 := 0

	for i := 1; i < len(arr); i++ {
		take := arr[i]
		if i > 1 {
			take += prev2
		}
		notTake := prev

		curri := max(take, notTake)
		prev2 = prev
		prev = curri
	}

	return prev
}

func robStreet(arr []int) int {
	arr1 := arr[1:]
	arr2 := arr[:len(arr)-1]

	return max(maxSteel(arr1), maxSteel(arr2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := []int{2, 1, 4, 9}
	m := []int{1, 5, 2, 1, 6}
	fmt.Println(robStreet(n)) // 10
	fmt.Println(robStreet(m)) // 11
}
