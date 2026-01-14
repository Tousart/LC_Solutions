package main

import (
	"fmt"
	"math"
)

func setZeroes(matrix [][]int) {
	for row := range matrix {
		for col := range matrix[0] {
			if matrix[row][col] == 0 {
				for currCol := range matrix[0] {
					if matrix[row][currCol] == 0 {
						matrix[row][currCol] = math.MaxInt64
					} else {
						matrix[row][currCol] = 0
					}
				}
				break
			}
		}
	}

	for row := range matrix {
		for col := range matrix[0] {
			if matrix[row][col] == math.MaxInt64 {
				for currRow := range matrix {
					matrix[currRow][col] = 0
				}
			}
		}
	}
}

func main() {
	fmt.Println(math.MaxInt64 < (2e31 - 1))
}
