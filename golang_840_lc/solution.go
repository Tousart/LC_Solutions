package main

import "fmt"

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}

	var res int

	for row := range len(grid) - 2 {
		for col := range len(grid[0]) - 2 {
			nums := make([]int, 10)
			cols := make([]int, 3)
			rows := make([]int, 3)
			breakFlag := false

			for i := row; i < row+3; i++ {
				sum := 0

				for j := col; j < col+3; j++ {
					if grid[i][j] > 9 || grid[i][j] < 1 || nums[grid[i][j]] == 1 {
						breakFlag = true
						break
					}
					nums[grid[i][j]] = 1
					sum += grid[i][j]
					cols[j-col] += grid[i][j]
				}

				if breakFlag {
					break
				}

				rows[i-row] += sum
			}

			if breakFlag {
				continue
			}

			currSum := grid[row][col+2] + grid[row+1][col+1] + grid[row+2][col]
			if currSum != (grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2]) {
				continue
			}

			for k := range 3 {
				if currSum != cols[k] || currSum != rows[k] {
					breakFlag = true
					break
				}
			}

			if breakFlag {
				continue
			}

			res++
		}
	}

	return res
}

func main() {
	grid := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	fmt.Println(numMagicSquaresInside(grid))

	grid1 := [][]int{
		{7, 0, 5},
		{2, 4, 6},
		{3, 8, 1},
	}
	fmt.Println(numMagicSquaresInside(grid1))

	grid2 := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	fmt.Println(numMagicSquaresInside(grid2))
}
