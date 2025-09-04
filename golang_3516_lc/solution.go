package main

import (
	"fmt"
	"math"
)

func findClosest(x int, y int, z int) int {
	xLen := math.Abs(float64(x - z))
	yLen := math.Abs(float64(y - z))
	if xLen < yLen {
		return 1
	} else if yLen < xLen {
		return 2
	}
	return 0
}

func main() {
	x, y, z := 2, 7, 4
	fmt.Println(findClosest(x, y, z))
}
