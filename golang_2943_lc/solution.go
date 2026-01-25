package main

import (
	"fmt"
	"sort"
)

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)
	res := min(getLen(hBars), getLen(vBars))
	return res * res
}

func getLen(bars []int) int {
	length, curr := 1, bars[0]-1
	res := 1
	for i := range len(bars) {
		if curr == bars[i]-1 {
			length++
		} else {
			res = max(res, length)
			length = 2
		}
		curr = bars[i]
	}
	return max(res, length)
}

func main() {
	n, m := 2, 1
	hBars := []int{2, 3}
	vBars := []int{2}
	fmt.Println(maximizeSquareHoleArea(n, m, hBars, vBars))

	n1, m1 := 14, 4
	hBars1 := []int{13}
	vBars1 := []int{3, 4, 5, 2}
	fmt.Println(maximizeSquareHoleArea(n1, m1, hBars1, vBars1))

	n2, m2 := 4, 40
	hBars2 := []int{5, 3, 2, 4}
	vBars2 := []int{36, 41, 6, 34, 33}
	fmt.Println(maximizeSquareHoleArea(n2, m2, hBars2, vBars2))
}
