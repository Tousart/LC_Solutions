package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findRoutes(n int) int {
	dp := make([]int, n)

	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	}

	dp[0] = 1
	dp[1] = 2
	dp[2] = 4

	for i := 3; i < n; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}

	return dp[n-1]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	writer.WriteString(strconv.Itoa(findRoutes(n)))
}
