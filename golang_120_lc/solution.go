package main

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := range len(triangle[i]) {
			if triangle[i+1][j+1] < triangle[i+1][j] {
				triangle[i][j] += triangle[i+1][j+1]
			} else {
				triangle[i][j] += triangle[i+1][j]
			}
		}
	}
	return triangle[0][0]
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}
