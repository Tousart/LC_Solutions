package main

import (
	"fmt"
	"math"
)

func countTrapezoids(points [][]int) int {
	var res int
	slopeEq := make(map[float64][][]float64)
	have := make(map[[4]float64]struct{})

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			x1, y1 := float64(points[i][0]), float64(points[i][1])
			x2, y2 := float64(points[j][0]), float64(points[j][1])

			var (
				tg float64
				eq float64
			)
			if x1 == x2 {
				tg = math.Pi / 2
				eq = x1
			} else if y1 == y2 {
				tg = 0
				eq = y1
			} else {
				tg = (y2 - y1) / (x2 - x1)
				eq = y1 - (x1 * tg)
			}

			if _, ok := slopeEq[tg]; !ok {
				slopeEq[tg] = [][]float64{{eq, x1, y1, x2, y2}}
				continue
			}

			for _, val := range slopeEq[tg] {
				if eq != val[0] {
					d1 := math.Sqrt(math.Pow(y2-y1, 2) + math.Pow(x2-x1, 2))
					d2 := math.Sqrt(math.Pow(val[4]-val[2], 2) + math.Pow(val[3]-val[1], 2))
					if d1 == d2 {
						coords := [][]float64{{x1, y1}, {x2, y2}, {val[1], val[2]}, {val[3], val[4]}}
						minX := coords[0][0]
						minY := coords[0][1]
						maxX := coords[0][0]
						maxY := coords[0][1]
						for k := 1; k < 4; k++ {
							if coords[k][0] <= minX {
								minX = coords[k][0]
								minY = min(minY, coords[k][1])
							}
							if coords[k][0] >= maxX {
								maxX = coords[k][0]
								maxY = max(maxY, coords[k][1])
							}
						}

						key := [4]float64{minX, minY, maxX, maxY}
						if _, ok := have[key]; ok {
							continue
						}
						have[key] = struct{}{}
					}
					res++
				}
			}

			slopeEq[tg] = append(slopeEq[tg], []float64{eq, x1, y1, x2, y2})
		}
	}

	return res
}

// a := fmt.Sprintf("(%d, %d) - (%d, %d)", x1, y1, x2, y2)
// b := fmt.Sprintf("(%.0f, %.0f) - (%.0f, %.0f)\n", val[1], val[2], val[3], val[4])
// fmt.Println(a)
// fmt.Println(b)
// fmt.Println()

func main() {
	points := [][]int{
		{-3, 2},
		{3, 0},
		{2, 3},
		{3, 2},
		{2, -3},
	}
	fmt.Println(countTrapezoids(points))

	points1 := [][]int{
		{-32, 12},
		{-32, -94},
		{-32, -15},
		{-30, 88},
	}
	fmt.Println(countTrapezoids(points1))

	points2 := [][]int{
		{71, -89},
		{-75, -89},
		{-9, 11},
		{-24, -89},
		{-51, -89},
		{-77, -89},
		{42, 11},
	}
	fmt.Println(countTrapezoids(points2))

	points3 := [][]int{
		{10, -66},
		{-36, 30},
		{86, 30},
		{10, 30},
		{-75, 19},
		{83, 19},
		{86, 19},
		{-39, 19},
	}
	fmt.Println(countTrapezoids(points3))
}
