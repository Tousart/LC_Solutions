package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func riverGo(river string) int {
	n := len(river)
	dp := make([]int, n*2+2)
	up := 1
	down := 1
	idx := 2

	for i := range n {
		if river[i] == 'B' {
			dp[idx], dp[idx+1] = dp[idx-2]+1, dp[idx-1]+1
			idx += 2
			continue
		}

		a := 0
		if (up == 1 && river[i] == 'L') || (up == 0 && river[i] == 'R') {
			a = 1
		}

		b := 0
		if (down == 0 && river[i] == 'R') || (down == 1 && river[i] == 'L') {
			b = 1
		}

		dp[idx], dp[idx+1] = dp[idx-2], dp[idx-1]

		if a == 1 && b == 1 {
			dp[idx], dp[idx+1] = dp[idx]+1, dp[idx+1]+1
			up = 1
			down = 0
		} else if a == 0 && b == 1 {
			up = 1
			down = 1
		} else if a == 1 && b == 0 {
			up = 0
			down = 0
		}

		idx += 2
	}

	if down == 1 {
		return dp[len(dp)-1] + 1
	}

	return dp[len(dp)-1]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')

	writer.WriteString(strconv.Itoa(riverGo(strings.TrimSpace(line))))
}
