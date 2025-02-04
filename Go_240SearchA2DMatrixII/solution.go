package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	var (
		row  = 0
		col  = len(matrix[0]) - 1
		rows = len(matrix) - 1
	)

	for col >= 0 && row <= rows {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}

	return false
}

func main() {
	var matrix = [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	fmt.Println(searchMatrix(matrix, 4))
}
