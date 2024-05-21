// https://www.youtube.com/watch?v=AxNNVECce8c&list=PLgUwDviBIf0rGlzIn_7rsaR2FQ5e6ZOL9&index=6&pp=iAQB

package main

import "fmt"

func printSequence(idx int, arr []int, temp []int) {
	if idx >= len(arr) {
		fmt.Println(temp)
		return
	}

	temp = append(temp, arr[idx])
	printSequence(idx+1, arr, temp)
	temp = temp[:len(temp)-1]
	printSequence(idx+1, arr, temp)
}

func main() {
	arr := []int{3, 1, 2}
	temp := []int{}
	printSequence(0, arr, temp)
}
