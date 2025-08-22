package main

import "fmt"

func minimumArea(grid [][]int) int {
	var (
		minRow int = -1
		minCol int
		maxRow int
		maxCol int
	)

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				if minRow == -1 {
					minRow = i
					minCol = j
					maxCol = j
				} else {
					minCol = min(minCol, j)
					maxCol = max(maxCol, j)
				}
				maxRow = i
			}
		}
	}

	return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{0, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(minimumArea(grid))
}
