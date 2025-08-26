package main

import (
	"fmt"
	"math"
)

func areaOfMaxDiagonal(dimensions [][]int) int {
	var (
		diag float64
		area float64
	)

	for i := range dimensions {
		a := float64(dimensions[i][0])
		b := float64(dimensions[i][1])
		newDiag := math.Pow(math.Pow(a, 2)+math.Pow(b, 2), 0.5)
		if newDiag > diag {
			diag = newDiag
			area = a * b
		} else if newDiag == diag && a*b > area {
			area = a * b
		}
	}

	return int(area)
}

func main() {
	dimensions := [][]int{
		{9, 3},
		{8, 6},
	}
	fmt.Println(areaOfMaxDiagonal(dimensions))
}
