package main

import "fmt"

type Queue []int

func (q *Queue) Enqueue(n int) {
	(*q) = append((*q), n)
}
func (q *Queue) Dequeue() int {
	ele := (*q)[0]
	(*q) = (*q)[1:]
	return ele
}
func (q *Queue) IsEmpty() bool {
	return len((*q)) == 0
}

func main() {
	adjList := graph(5, [][]int{{1, 2}, {2, 4}, {1, 3}, {3, 4}, {2, 5}, {4, 5}})
	fmt.Println(bfs(1, adjList))
}

func graph(n int, edges [][]int) [][]int {
	adjList := make([][]int, n+1)

	for i := 0; i < len(edges); i++ {
		a := edges[i][0]
		b := edges[i][1]

		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}

	return adjList
}

func bfs(startIdx int, adjList [][]int) []int {
	q := Queue{}
	visited := make([]int, len(adjList))
	nodes := []int{}

	q.Enqueue(startIdx)
	visited[startIdx] = 1

	for !q.IsEmpty() {
		idx := q.Dequeue()
		nodes = append(nodes, idx)
		for _, val := range adjList[idx] {
			if visited[val] != 1 {
				q.Enqueue(val)
				visited[val] = 1
			}
		}
	}

	return nodes
}
