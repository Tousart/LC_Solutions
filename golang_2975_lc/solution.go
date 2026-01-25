package main

import (
	"fmt"
	"math"
)

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	mod := 1000000007

	if m == n {
		return ((m - 1) * (n - 1)) % mod
	}

	res := -1
	lenMap := make(map[int]struct{})
	lenMap[m-1] = struct{}{}

	for i := range len(hFences) {
		lenMap[m-hFences[i]] = struct{}{}
		lenMap[hFences[i]-1] = struct{}{}

		for j := i + 1; j < len(hFences); j++ {
			lenMap[int(math.Abs(float64(hFences[j]-hFences[i])))] = struct{}{}
		}
	}

	if _, ok := lenMap[n-1]; ok {
		return ((n - 1) * (n - 1)) % mod
	}

	for i := range len(vFences) {
		if _, ok := lenMap[n-vFences[i]]; ok {
			res = max(res, (n-vFences[i])*(n-vFences[i]))
		}

		if _, ok := lenMap[vFences[i]-1]; ok {
			res = max(res, (vFences[i]-1)*(vFences[i]-1))
		}

		for j := i + 1; j < len(vFences); j++ {
			length := int(math.Abs(float64(vFences[j] - vFences[i])))
			if _, ok := lenMap[length]; ok {
				res = max(res, length*length)
			}
		}
	}

	return res % mod
}

func main() {
	m := 4
	n := 3
	hFenches := []int{2, 3}
	vFenches := []int{2}
	fmt.Println(maximizeSquareArea(m, n, hFenches, vFenches))

	m1 := 4
	n1 := 5
	hFenches1 := []int{2}
	vFenches1 := []int{4}
	fmt.Println(maximizeSquareArea(m1, n1, hFenches1, vFenches1))

	m2 := 3
	n2 := 9
	hFenches2 := []int{2}
	vFenches2 := []int{8, 6, 5, 4}
	fmt.Println(maximizeSquareArea(m2, n2, hFenches2, vFenches2))
}
