package main

import "fmt"

func maxSideLength(mat [][]int, threshold int) int {
	var res int
	m, n := len(mat), len(mat[0])
	prefix := make([][]int, m+1)
	prefix[0] = make([]int, n+1)

	for row := 1; row <= m; row++ {
		prefix[row] = make([]int, n+1)
		for col := 1; col <= n; col++ {
			prefix[row][col] = mat[row-1][col-1] + prefix[row-1][col] + prefix[row][col-1] - prefix[row-1][col-1]
		}
	}

	getArea := func(row, col, k int) int {
		return prefix[row+k][col+k] + prefix[row][col] - prefix[row+k][col] - prefix[row][col+k]
	}

	board := min(m, n)
	for row := range m {
		for col := range n {
			for k := res + 1; k <= board; k++ {
				if (row+k <= m) && (col+k <= n) && getArea(row, col, k) <= threshold {
					res = k
				} else {
					break
				}
			}
		}
	}

	return res
}

func main() {
	mat := [][]int{
		{1, 1, 3, 2, 4, 3, 2},
		{1, 1, 3, 2, 4, 3, 2},
		{1, 1, 3, 2, 4, 3, 2},
	}
	threshold := 4
	fmt.Println(maxSideLength(mat, threshold))

	mat1 := [][]int{
		{0},
	}
	threshold1 := 0
	fmt.Println(maxSideLength(mat1, threshold1))

	mat2 := [][]int{
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
	}
	threshold2 := 1
	fmt.Println(maxSideLength(mat2, threshold2))

	mat3 := [][]int{
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	threshold3 := 6
	fmt.Println(maxSideLength(mat3, threshold3))
}
