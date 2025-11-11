package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func maxRoute(route [][]rune) int {
	dp := make([]int, 3)
	for i := range 3 {
		switch route[0][i] {
		case 'C':
			dp[i] = 1
		case '.':
			dp[i] = 0
		case 'W':
			dp[i] = -1
		}
	}

	if dp[0] == -1 && dp[1] == -1 && dp[2] == -1 {
		return 0
	}

	temp := make([]int, 3)

	for row := 1; row < len(route); row++ {
		left := max(dp[0], dp[1])
		right := max(dp[1], dp[2])
		center := max(left, right)

		if center != -1 {
			copy(temp, dp)

			switch route[row][1] {
			case 'C':
				dp[1] = 1 + center
			case '.':
				dp[1] = center
			case 'W':
				dp[1] = -1
			}
		} else {
			return max(temp[0], temp[1], temp[2])
		}

		if left != -1 {
			switch route[row][0] {
			case 'C':
				dp[0] = 1 + left
			case '.':
				dp[0] = left
			case 'W':
				dp[0] = -1
			}
		} else {
			dp[0] = -1
		}

		if right != -1 {
			switch route[row][2] {
			case 'C':
				dp[2] = 1 + right
			case '.':
				dp[2] = right
			case 'W':
				dp[2] = -1
			}
		} else {
			dp[2] = -1
		}
	}

	return max(dp[0], dp[1], dp[2])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line1))

	route := make([][]rune, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		route[i] = []rune(strings.TrimSpace(line))
	}

	writer.WriteString(strconv.Itoa(maxRoute(route)))
}
