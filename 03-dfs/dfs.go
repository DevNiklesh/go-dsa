package main

import "fmt"

func main() {
	nodes := []int{}
	adjList := graph(5, [][]int{{1, 2}, {2, 4}, {1, 3}, {3, 4}, {2, 5}, {4, 5}})
	visited := make([]int, len(adjList)+1)

	dfs(1, adjList, &nodes, &visited)
	fmt.Println(nodes)
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

func dfs(node int, adjList [][]int, list *[]int, visited *[]int) {
	(*visited)[node] = 1
	(*list) = append((*list), node)

	for _, val := range adjList[node] {
		if (*visited)[val] != 1 {
			dfs(val, adjList, list, visited)
		}
	}
}
