package main

import "fmt"

func largestMagicSquare(grid [][]int) int {
	res := 1
	m, n := len(grid), len(grid[0])

	rows := make([][]int, m)
	cols := make([][]int, m)
	for row := range m {
		rows[row] = make([]int, n)
		cols[row] = make([]int, n)
		for col := range n {
			rows[row][col] = grid[row][col]
			cols[row][col] = grid[row][col]

			if col > 0 {
				rows[row][col] += rows[row][col-1]
			}

			if row > 0 {
				cols[row][col] += cols[row-1][col]
			}
		}
	}

	for row := range m {
		for col := range n {
			r, c := row, col
			for r < m && c < n {
				magic := true

				sumRows := rows[row][c]
				if col > 0 {
					sumRows -= rows[row][col-1]
				}

				sumCols := cols[r][col]
				if row > 0 {
					sumCols -= cols[row-1][col]
				}

				if sumCols != sumRows {
					r++
					c++
					continue
				}

				diagL := grid[row][col]
				idx := col + 1
				for k := row + 1; k <= r; k++ {
					if col > 0 {
						if (rows[k][c] - rows[k][col-1]) != sumRows {
							magic = false
							break
						}
					} else if rows[k][c] != sumRows {
						magic = false
						break
					}
					diagL += grid[k][idx]
					idx++
				}

				if !magic || diagL != sumRows {
					r++
					c++
					continue
				}

				diagR := grid[r][col]
				idx = r - 1
				for k := col + 1; k <= c; k++ {
					if row > 0 {
						if (cols[r][k] - cols[row-1][k]) != sumCols {
							magic = false
							break
						}
					} else if cols[r][k] != sumCols {
						magic = false
						break
					}
					diagR += grid[idx][k]
					idx--
				}

				if !magic || diagR != sumCols || diagL != diagR {
					r++
					c++
					continue
				}

				res = max(res, r-row+1)

				r++
				c++
			}
		}
	}

	return res
}

func main() {
	grid := [][]int{
		{7, 1, 4, 5, 6},
		{2, 5, 1, 6, 4},
		{1, 5, 4, 3, 2},
		{1, 2, 7, 3, 4},
	}
	fmt.Println(largestMagicSquare(grid))
}

// func largestMagicSquare(grid [][]int) int {
// 	res := 1
// 	m, n := len(grid), len(grid[0])

// 	for row := range m {
// 		for col := range n {
// 			rows := make([]int, 0)
// 			cols := make([]int, 0)
// 			diagL := 0

// 			i, j := row, col
// 			for i < m && j < n {
// 				// diag left
// 				diagL += grid[i][j]

// 				// rows
// 				for k := row; k < i; k++ {
// 					rows[k-row] += grid[k][j]
// 				}
// 				rows = append(rows, 0)
// 				for k := col; k <= j; k++ {
// 					rows[len(rows)-1] += grid[i][k]
// 				}

// 				// cols
// 				for k := col; k < j; k++ {
// 					cols[k-col] += grid[i][k]
// 				}
// 				cols = append(cols, 0)
// 				for k := row; k <= i; k++ {
// 					cols[len(cols)-1] += grid[k][j]
// 				}

// 				// diag right
// 				diagR := 0
// 				r, c := row, j
// 				for r <= i && c >= col {
// 					diagR += grid[r][c]
// 					r++
// 					c--
// 				}

// 				// chek sums
// 				if diagR == diagL {
// 					idx := 0
// 					for idx < len(rows) {
// 						if rows[idx] != diagL || cols[idx] != diagL {
// 							break
// 						}
// 						idx++
// 					}

// 					if idx == len(rows) {
// 						res = max(res, i-row+1)
// 					}
// 				}

// 				i++
// 				j++
// 			}
// 		}
// 	}

// 	return res
// }
