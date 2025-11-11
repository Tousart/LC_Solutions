package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findMaxDiff(row, col int, matrix []string) int {
	maxPlus := -1001
	minMinus := 1001
	rows := make(map[int][]int)
	cols := make(map[int][]int)

	for i := range row {
		val := 0
		for j := range col {
			if matrix[i][j] == '-' {
				val--
			} else {
				val++
			}
		}

		if val >= maxPlus {
			if _, ok := rows[val]; !ok {
				rows[val] = []int{i}
			} else {
				rows[val] = append(rows[val], i)
			}
			maxPlus = val
		}
	}

	for j := range col {
		val := 0
		for i := range row {
			if matrix[i][j] == '+' {
				val++
			} else {
				val--
			}
		}

		if val <= minMinus {
			if _, ok := cols[val]; !ok {
				cols[val] = []int{j}
			} else {
				cols[val] = append(cols[val], j)
			}
			minMinus = val
		}
	}

	res := maxPlus - minMinus - 2
	for _, rowIdx := range rows[maxPlus] {
		for _, colIdx := range cols[minMinus] {
			if matrix[rowIdx][colIdx] != '?' {
				return res + 2
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
	rowCol := strings.Split(strings.TrimSpace(line1), " ")
	row, _ := strconv.Atoi(rowCol[0])
	col, _ := strconv.Atoi(rowCol[1])
	matrix := make([]string, row)
	for i := range row {
		line, _ := reader.ReadString('\n')
		matrix[i] = strings.TrimSpace(line)
	}

	writer.WriteString(strconv.Itoa(findMaxDiff(row, col, matrix)))
}
