package main

import (
	"fmt"
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	var res float64
	for i := range len(points) - 2 {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				res = max(res, math.Abs(float64((points[j][0]-points[i][0])*(points[k][1]-points[i][1])-(points[k][0]-points[i][0])*(points[j][1]-points[i][1])))/2)
			}
		}
	}
	return res
}

func main() {
	points := [][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{0, 2},
		{2, 0},
	}
	fmt.Println(largestTriangleArea(points))
}
