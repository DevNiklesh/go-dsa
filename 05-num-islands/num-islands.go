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

func bfs(r, c int, grid [][]byte, visited *[][]int) {
	q := &Queue{}

	(*visited)[r][c] = 1
	q.Enqueue([]int{r, c})

	for !q.IsEmpty() {
		node := q.Dequeue()
		row, col := node[0], node[1]

		if col+1 < len(grid[row]) && grid[row][col+1] == '1' {
			if (*visited)[row][col+1] != 1 {
				(*visited)[row][col+1] = 1
				q.Enqueue([]int{row, col + 1})
			}
		}
		if col-1 >= 0 && grid[row][col-1] == '1' {
			if (*visited)[row][col-1] != 1 {
				(*visited)[row][col-1] = 1
				q.Enqueue([]int{row, col - 1})
			}
		}
		if row-1 >= 0 && grid[row-1][col] == '1' {
			if (*visited)[row-1][col] != 1 {
				(*visited)[row-1][col] = 1
				q.Enqueue([]int{row - 1, col})
			}
		}
		if row+1 < len(grid) && grid[row+1][col] == '1' {
			if (*visited)[row+1][col] != 1 {
				(*visited)[row+1][col] = 1
				q.Enqueue([]int{row + 1, col})
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[i]))
	}

	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '1' && visited[r][c] == 0 { // not visited
				count++
				bfs(r, c, grid, &visited)
			}
		}
	}

	return count
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))

	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}
