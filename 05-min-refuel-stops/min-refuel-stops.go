package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Push(n interface{}) {
	(*h) = append((*h), n.(int))
}
func (h *MaxHeap) Pop() interface{} {
	original := *h
	l := len(original)
	x := original[l-1]
	*h = (*h)[:l-1]

	return x
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	maxHeap := &MaxHeap{}

	maxDistance := startFuel
	i := 0
	stops := 0
	for maxDistance < target {
		if i < len(stations) && stations[i][0] <= maxDistance {
			heap.Push(maxHeap, stations[i][1])
			i++
		} else if maxHeap.Len() == 0 {
			return -1
		} else {
			maxDistance = maxDistance + heap.Pop(maxHeap).(int)
			stops++
		}
		fmt.Println("maxDistance", maxDistance, "maxHeap", maxHeap, "stops", stops)
	}

	return stops
}

func main() {
	fmt.Println(minRefuelStops(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}))
}
