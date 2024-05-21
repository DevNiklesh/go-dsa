package main

import "fmt"

func main() {
	fmt.Println(graph(5, [][]int{{1, 2}, {2, 4}, {1, 3}, {3, 4}, {2, 5}, {4, 5}}))
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
