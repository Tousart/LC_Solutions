package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func countOfStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for basis := 1; basis <= n; basis++ {
		for i := n; i >= basis; i-- {
			dp[i] += dp[i-basis]
		}
	}
	return dp[n]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	writer.WriteString(strconv.Itoa(countOfStairs(n)))
}
