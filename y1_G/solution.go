package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func yesOrNo(n, m int, matrix []string) string {
	for row := range n {
		for col := range m {
			if matrix[row][col] == '.' {
				continue
			}

			current := matrix[row][col]

			if col+5 <= m {
				right := col
				count := 0
				for right < m {
					if matrix[row][right] == current {
						count++
					} else if matrix[row][right] != current {
						break
					}

					right++

					if count == 5 {
						return "Yes"
					}
				}
			}

			if row+5 <= n {
				down := row
				count := 0
				for down < n {
					if matrix[down][col] == current {
						count++
					} else if matrix[down][col] != current {
						break
					}

					down++

					if count == 5 {
						return "Yes"
					}
				}

				if col+5 <= m {
					diagonalRow := row
					diagonalCol := col
					count := 0
					for (diagonalRow < n) && (diagonalCol < m) {
						if matrix[diagonalRow][diagonalCol] == current {
							count++
						} else if matrix[diagonalRow][diagonalCol] != current {
							break
						}

						diagonalRow++
						diagonalCol++

						if count == 5 {
							return "Yes"
						}
					}
				}

				if col-5 >= -1 {
					diagonalRow := row
					diagonalCol := col
					count := 0
					for (diagonalRow < n) && (diagonalCol >= 0) {
						if matrix[diagonalRow][diagonalCol] == current {
							count++
						} else if matrix[diagonalRow][diagonalCol] != current {
							break
						}

						diagonalRow++
						diagonalCol--

						if count == 5 {
							return "Yes"
						}
					}
				}
			}
		}
	}
	return "No"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	nm := strings.Split(strings.TrimSpace(line1), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	matrix := make([]string, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		matrix[i] = strings.TrimSpace(line)
	}

	writer.WriteString(yesOrNo(n, m, matrix))
}
