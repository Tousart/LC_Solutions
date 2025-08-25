package main

import (
	"fmt"
)

func drain(tanks [][3]int) int {
	var res int
	n := len(tanks)
	trash := make([]int, n)
	right := make([]int, n)
	left := make([]int, n)

	// right
	for i := 1; i < n; i++ {
		if tanks[i-1][1] > tanks[i][2] {
			right[i] = tanks[i][2]
			trash[i] = tanks[i-1][1] - tanks[i][2]
		} else {
			right[i] = tanks[i-1][1]
		}
	}

	for i := n - 2; i > 0; i-- {
		right[i] += min(right[i+1], trash[i])
	}

	// left
	for i := n - 2; i >= 0; i-- {
		if tanks[i+1][0] > tanks[i][2] {
			left[i] = tanks[i][2]
			trash[i] = tanks[i+1][0] - tanks[i][2]
		} else {
			left[i] = tanks[i+1][0]
		}
	}

	for i := 1; i < n-1; i++ {
		left[i] += min(left[i-1], trash[i])
	}

	// answer
	res = max(tanks[0][2]+right[1], tanks[n-1][2]+left[n-2])
	for i := 1; i < n-1; i++ {
		res = max(res, tanks[i][2]+right[i+1]+left[i-1])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)
	tanks := make([][3]int, n)
	for i := range n {
		fmt.Scan(&tanks[i][0], &tanks[i][1], &tanks[i][2])
	}
	fmt.Println(drain(tanks))
}
