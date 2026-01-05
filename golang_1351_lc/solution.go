package main

import "fmt"

func countNegatives(grid [][]int) int {
	var res int

	row := 0
	col := len(grid[0]) - 1
	for row < len(grid) && col >= 0 {
		if grid[row][col] < 0 {
			res += len(grid) - row
			col--
		} else {
			row++
		}
	}

	return res
}

func main() {
	grid := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	fmt.Println(countNegatives(grid))
}
