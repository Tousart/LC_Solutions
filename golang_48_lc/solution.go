package main

func rotate(matrix [][]int) {
	idx := 0
	for row := range matrix {
		for col := idx; col < len(matrix[0]); col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}

		l, r := 0, len(matrix[0])-1
		for l < r {
			matrix[row][l], matrix[row][r] = matrix[row][r], matrix[row][l]
			l++
			r--
		}

		idx++
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
}
