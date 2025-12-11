package main

import "fmt"

func countCoveredBuildings(n int, buildings [][]int) int {
	var res int

	xLimits := make([][]int, n+1)
	yLimits := make([][]int, n+1)

	for _, coords := range buildings {
		x, y := coords[0], coords[1]

		if xLimits[x] == nil {
			xLimits[x] = []int{y, y}
		} else {
			xLimits[x][0], xLimits[x][1] = min(xLimits[x][0], y), max(xLimits[x][1], y)
		}

		if yLimits[y] == nil {
			yLimits[y] = []int{x, x}
		} else {
			yLimits[y][0], yLimits[y][1] = min(yLimits[y][0], x), max(yLimits[y][1], x)
		}
	}

	for _, coords := range buildings {
		x, y := coords[0], coords[1]

		if y > xLimits[x][0] && y < xLimits[x][1] && x > yLimits[y][0] && x < yLimits[y][1] {
			res++
		}
	}

	return res
}

func main() {
	n := 3
	buildings := [][]int{
		{1, 2},
		{2, 2},
		{3, 2},
		{2, 1},
		{2, 3},
	}
	fmt.Println(countCoveredBuildings(n, buildings))
}
