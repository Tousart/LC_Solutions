package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	var (
		count int
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				bfs(i, j, grid)
				//dfs(i, j, grid)
			}
		}
	}

	return count
}

func dfs(n, m int, grid [][]byte) {
	if n == len(grid) || n < 0 || m == len(grid[0]) || m < 0 || grid[n][m] != '1' {
		return
	}

	grid[n][m] = '2'
	dfs(n+1, m, grid)
	dfs(n-1, m, grid)
	dfs(n, m+1, grid)
	dfs(n, m-1, grid)
}

func bfs(n, m int, grid [][]byte) {
	q := [][2]int{{n, m}}
	offsets := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	grid[n][m] = '2'
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		for _, offset := range offsets {
			row := current[0] + offset[0]
			col := current[1] + offset[1]

			if row == len(grid) || row < 0 || col == len(grid[0]) || col < 0 || grid[row][col] != '1' {
				continue
			}

			q = append(q, [2]int{row, col})
			grid[row][col] = '2'
		}
	}
}

func main() {
	grid := [][]byte{
		[]byte("1111"),
		[]byte("1111"),
	}
	fmt.Println(numIslands(grid))
}
