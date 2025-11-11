package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxSequence(n, m int, matrix [][]int) int {
	res := 1
	dp := make([][]int, n)
	vrc := [][3]int{}
	for row := range n {
		dp[row] = make([]int, m)
		for col := range m {
			dp[row][col] = 1
			vrc = append(vrc, [3]int{matrix[row][col], row, col})
		}
	}

	sort.Slice(vrc, func(i, j int) bool {
		return vrc[i][0] < vrc[j][0]
	})

	coords := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for _, curr := range vrc {
		value, row, col := curr[0], curr[1], curr[2]
		for _, coord := range coords {
			x, y := row+coord[0], col+coord[1]
			if (x >= 0) && (x < n) && (y >= 0) && (y < m) {
				if value-matrix[x][y] == 1 {
					dp[row][col] = max(dp[row][col], dp[x][y]+1)
					res = max(res, dp[row][col])
				}
			}
		}
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	nm := strings.Split(strings.TrimSpace(line1), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	matrix := make([][]int, n)
	for i := range n {
		matrix[i] = make([]int, m)
		line, _ := reader.ReadString('\n')
		input := strings.Split(strings.TrimSpace(line), " ")
		for j := range input {
			val, _ := strconv.Atoi(input[j])
			matrix[i][j] = val
		}
	}

	writer.WriteString(strconv.Itoa(maxSequence(n, m, matrix)))
}
