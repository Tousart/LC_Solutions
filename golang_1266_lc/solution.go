package main

import (
	"fmt"
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	var res int
	x, y := points[0][0], points[0][1]

	for i := 1; i < len(points); i++ {
		if x == points[i][0] {
			res += int(math.Abs(float64(y - points[i][1])))
		} else if y == points[i][1] {
			res += int(math.Abs(float64(x - points[i][0])))
		} else {
			res += max(int(math.Abs(float64(x-points[i][0]))), int(math.Abs(float64(y-points[i][1]))))
		}
		x, y = points[i][0], points[i][1]
	}

	return res
}

func main() {
	points := [][]int{
		{1, 1},
		{3, 4},
		{-1, 0},
	}
	fmt.Println(minTimeToVisitAllPoints(points))
}
