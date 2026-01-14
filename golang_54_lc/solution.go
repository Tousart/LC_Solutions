package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	res := make([]int, len(matrix)*len(matrix[0]))
	idx := 0
	l, r, u, d := 0, len(matrix[0])-1, 0, len(matrix)-1

	for idx < len(res) {
		for i := l; i <= r; i++ {
			res[idx] = matrix[u][i]
			idx++
		}
		u++

		if idx == len(res) {
			break
		}

		for i := u; i <= d; i++ {
			res[idx] = matrix[i][r]
			idx++
		}
		r--

		if idx == len(res) {
			break
		}

		for i := r; i >= l; i-- {
			res[idx] = matrix[d][i]
			idx++
		}
		d--

		if idx == len(res) {
			break
		}

		for i := d; i >= u; i-- {
			res[idx] = matrix[i][l]
			idx++
		}
		l++
	}

	return res
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}
