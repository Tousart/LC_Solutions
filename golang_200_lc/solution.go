package main

import "fmt"

func numIslands(grid [][]byte) int {
	var res int

	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == '1' {
				res++
				// bfs(grid, row, col)
				dfs(grid, row, col)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, row, col int) {
	if (row >= 0 && row < len(grid)) &&
		(col >= 0 && col < len(grid[0])) &&
		(grid[row][col] == '1') {
		grid[row][col] = '2'
	} else {
		return
	}

	dfs(grid, row+1, col)
	dfs(grid, row-1, col)
	dfs(grid, row, col+1)
	dfs(grid, row, col-1)
}

func bfs(grid [][]byte, row, col int) {
	coords := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][]int{{row, col}}
	grid[row][col] = '2'

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, coord := range coords {
			newX := x + coord[0]
			newY := y + coord[1]
			if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) {
				if grid[newX][newY] == '1' {
					queue = append(queue, []int{newX, newY})
					grid[newX][newY] = '2'
				}
			}
		}
	}
}

func main() {
	grid := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}
	fmt.Println(numIslands(grid))
}
