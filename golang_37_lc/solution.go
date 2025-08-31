package main

import "fmt"

func solveSudoku(board [][]byte) {
	addNum(board)
}

func addNum(board [][]byte) bool {
	for row := range 9 {
		for col := range 9 {
			if board[row][col] == '.' {
				for num := byte('1'); num <= '9'; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						if addNum(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, num byte) bool {
	for i := range 9 {
		if board[row][i] == num ||
			board[i][col] == num ||
			board[(row/3)*3+i/3][(col/3)*3+i%3] == num {
			return false
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
	solveSudoku(board)
	fmt.Println(board)
}
