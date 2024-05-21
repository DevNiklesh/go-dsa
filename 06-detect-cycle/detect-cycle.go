package main

import "fmt"

type Queue [][]int

func (q *Queue) Enqueue(n []int) {
	(*q) = append((*q), n)
}
func (q *Queue) Dequeue() []int {
	ele := (*q)[0]
	(*q) = (*q)[1:]
	return ele
}
func (q *Queue) IsEmpty() bool {
	return len((*q)) == 0
}

func isCycle(adjList [][]int) bool {
	visited := make([]int, len(adjList))

	for i := range adjList {
		if visited[i] != 1 {
			if detectCycle(i, adjList, &visited) {
				return true
			}
		}
	}

	return false
}

func main() {
	adjList := [][]int{}
	fmt.Println(isCycle(adjList))
}

func detectCycle(src int, adjList [][]int, visited *[]int) bool {
	q := &Queue{}

	(*visited)[src] = 1
	q.Enqueue([]int{src, -1})

	for !q.IsEmpty() {
		pair := q.Dequeue()
		node, parent := pair[0], pair[1]

		for _, adjNode := range adjList[node] {
			if (*visited)[adjNode] != 1 {
				q.Enqueue([]int{adjNode, node})
				(*visited)[adjNode] = 1
			} else if parent != adjNode {
				return true
			}
		}
	}

	return false
}
