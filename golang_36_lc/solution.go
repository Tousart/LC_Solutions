package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := make([][]int, 9)
	cols := make([][]int, 9)
	square := make([][]int, 9)

	for row := range 9 {
		rows[row] = make([]int, 10)
		for col := range 9 {
			if board[row][col] != '.' {
				if cols[col] == nil {
					cols[col] = make([]int, 10)
				}

				val := int(board[row][col] - '0')

				if rows[row][val] != 0 {
					return false
				}

				if cols[col][val] != 0 {
					return false
				}

				ind := (row/3)*3 + col/3
				if square[ind] == nil {
					square[ind] = make([]int, 10)
				} else if square[ind][val] != 0 {
					return false
				}

				rows[row][val] = val
				cols[col][val] = val
				square[ind][val] = val
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	fmt.Println(isValidSudoku(board))
}
