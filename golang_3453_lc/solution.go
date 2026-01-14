package main

import (
	"fmt"
	"math"
)

func separateSquares(squares [][]int) float64 {
	low, high, totalArea := 0.0, 1e9, 0.0

	for _, square := range squares {
		high = max(high, float64(square[1]))
		totalArea += math.Pow(float64(square[2]), 2)
	}

	minDiff := 1e-5

	for math.Abs(high-low) > minDiff {
		mid := (low + high) / 2.0

		if aboveTotalArea(squares, mid, totalArea) {
			low = mid
		} else {
			high = mid
		}
	}

	return low
}

func aboveTotalArea(squares [][]int, y, totalArea float64) bool {
	var above float64

	for _, square := range squares {
		up := float64(square[1] + square[2])
		if y < up {
			above += float64(square[2]) * min(float64(square[1]+square[2])-y, float64(square[2]))
		}
	}

	return above > totalArea/2.0
}

func main() {
	squares := [][]int{
		{0, 0, 2},
		{1, 1, 1},
	}
	fmt.Println(separateSquares(squares))

	squares1 := [][]int{
		{2, 5, 3},
		{8, 12, 4},
	}
	fmt.Println(separateSquares(squares1))
}

// func separateSquares(squares [][]int) float64 {
// 	low, high := 0.0, 2e9

// 	for range 60 {
// 		mid := (low + high) / 2.0
// 		diff := aboveBelowArea(squares, mid)

// 		if diff <= 0 {
// 			high = mid
// 		} else {
// 			low = mid
// 		}
// 	}

// 	return low
// }

// func aboveBelowArea(squares [][]int, y float64) float64 {
// 	var (
// 		above float64
// 		below float64
// 	)

// 	for _, square := range squares {
// 		down := float64(square[1])
// 		up := float64(square[2]) + down

// 		if y >= up {
// 			below += math.Pow(float64(square[2]), 2)
// 		} else if y <= down {
// 			above += math.Pow(float64(square[2]), 2)
// 		} else {
// 			edge := up - down
// 			above += edge * (up - y)
// 			below += edge * (y - down)
// 		}
// 	}

// 	return above - below
// }
