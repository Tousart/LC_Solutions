package main

import (
	"fmt"
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
	var res int

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		minHigh := math.MaxInt
		for j := i - 1; j >= 0; j-- {
			if (points[j][1] >= points[i][1]) && (minHigh > points[j][1]) {
				res++
				minHigh = points[j][1]
			}
		}
	}

	return res
}

func main() {
	points := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	fmt.Println(numberOfPairs(points))

	points1 := [][]int{
		{0, 0},
		{0, 3},
	}
	fmt.Println(numberOfPairs(points1))
}
