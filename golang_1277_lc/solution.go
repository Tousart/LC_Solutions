package main

import "fmt"

func countSquares(matrix [][]int) int {
	res := 0
	dp := make([][]int, len(matrix))
	for row := range matrix {
		dp[row] = make([]int, len(matrix[0]))
		for col := range matrix[row] {
			if matrix[row][col] == 1 {
				dp[row][col] = 1
				if row > 0 && col > 0 {
					dp[row][col] += min(dp[row-1][col], min(dp[row][col-1], dp[row-1][col-1]))
				}
				res += dp[row][col]
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}
	fmt.Println(countSquares(matrix))

	matrix1 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 1, 0},
		{1, 1, 1},
		{1, 1, 0},
	}
	fmt.Println(countSquares(matrix1))
}
