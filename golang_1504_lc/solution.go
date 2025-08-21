package main

import "fmt"

func numSubmat(mat [][]int) int {
	height := make([]int, len(mat[0]))
	n := len(mat[0])
	res := 0
	for i := range mat {
		for j := range n {
			if mat[i][j] == 0 {
				height[j] = 0
			} else {
				height[j]++
			}
		}

		stack := []int{}
		sums := make([]int, n)
		for j := range n {
			for len(stack) > 0 && height[j] <= height[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				indLastSum := stack[len(stack)-1]
				sums[j] = height[j]*(j-indLastSum) + sums[indLastSum]
			} else {
				sums[j] = height[j] * (j + 1)
			}

			stack = append(stack, j)
			res += sums[j]
		}
	}

	return res
}

func main() {
	mat := [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 0},
	}

	fmt.Println(numSubmat(mat))

	mat1 := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 1, 1, 0},
	}

	fmt.Println(numSubmat(mat1))
}
