package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] {
				if dfs(board, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, row, col, ind int) bool {
	if row < 0 || row == len(board) || col < 0 || col == len(board[0]) {
		return false
	}

	letter := board[row][col]
	if letter == word[ind] {
		if ind == len(word)-1 {
			return true
		}
		board[row][col] = '*'
	} else {
		return false
	}

	newInd := ind + 1
	found := dfs(board, word, row+1, col, newInd) ||
		dfs(board, word, row, col+1, newInd) ||
		dfs(board, word, row-1, col, newInd) ||
		dfs(board, word, row, col-1, newInd)

	board[row][col] = letter
	return found
}

func main() {
	board := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))

	board1 := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}
	word1 := "SEE"
	fmt.Println(exist(board1, word1))
}
