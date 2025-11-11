package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func sumOfMoney(days []int, n int) int {
	var res int
	prefix := make([]int, n)

	for i := range n {
		if i-1 >= 0 {
			prefix[i] += prefix[i-1]
		}

		if days[i] <= 1 {
			continue
		}

		end := i + days[i]

		if end < n {
			prefix[end]--
		}

		if i+1 < n {
			prefix[i+1]++
		}
	}

	for i := range prefix {
		res += days[i] * prefix[i]
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line1))

	line2, _ := reader.ReadString('\n')
	input := strings.Split(strings.TrimSpace(line2), " ")
	days := make([]int, n)
	for i := range n {
		elem, _ := strconv.Atoi(input[i])
		days[i] = elem
	}

	writer.WriteString(strconv.Itoa(sumOfMoney(days, n)))
}
