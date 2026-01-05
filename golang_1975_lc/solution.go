package main

import (
	"fmt"
	"math"
)

func maxMatrixSum(matrix [][]int) int64 {
	var (
		res           int
		countNegative int
		minNum        int = math.MaxInt64
	)

	for row := range matrix {
		for col := range matrix[0] {
			num := matrix[row][col]
			if num < 0 {
				countNegative++
				num *= -1
			}
			res += num
			minNum = min(minNum, num)
		}
	}

	if countNegative%2 == 0 {
		return int64(res)
	}

	return int64(res - 2*minNum)
}

func main() {
	matrix := [][]int{
		{1, -1},
		{-1, 1},
	}
	fmt.Println(maxMatrixSum(matrix))
}
