package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(n interface{}) {
	*h = append(*h, n.(int))
}
func (h *MinHeap) Pop() interface{} {
	original := *h
	l := len(*h)
	x := original[l-1]
	*h = original[:l-1]

	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	mh := &MinHeap{}
	heap.Init(mh)

	currBuildingIdx := 0
	availableBricks := bricks
	for currBuildingIdx = 0; currBuildingIdx < len(heights)-1; currBuildingIdx++ {

		if heights[currBuildingIdx] >= heights[currBuildingIdx+1] {
			continue
		}

		heap.Push(mh, heights[currBuildingIdx+1]-heights[currBuildingIdx])

		if mh.Len() > ladders {
			heightDiff := heap.Pop(mh)
			availableBricks -= heightDiff.(int)
		}

		if availableBricks < 0 {
			break
		}
	}

	return currBuildingIdx
}

func main() {
	fmt.Println(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
	fmt.Println(furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
	fmt.Println(furthestBuilding([]int{14, 3, 19, 3}, 17, 0))
}
