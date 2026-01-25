package main

import "fmt"

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	var side int
	n := len(bottomLeft)

	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			height := min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])
			width := min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])
			if height > 0 && width > 0 {
				side = max(side, min(height, width))
			}
		}
	}

	return int64(side * side)
}

func main() {
	bottomLeft := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
	}

	topRight := [][]int{
		{5, 5},
		{5, 7},
		{5, 9},
	}

	fmt.Println(largestSquareArea(bottomLeft, topRight))
}
