package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func whatSum(n, k int) int {
	if k == 0 {
		return n
	}

	remainders := []int{}
	exists := make(map[int]bool)
	n += n % 10
	lastNum := n % 10
	k--
	sum := 0
	for !exists[lastNum] {
		remainders = append(remainders, lastNum)
		exists[lastNum] = true
		sum += lastNum
		lastNum = (lastNum * 2) % 10
	}

	if len(remainders) > k {
		for i := range k {
			n += remainders[i]
		}
		return n
	}

	n += sum * (k / len(remainders))

	k = k % len(remainders)
	if k != 0 {
		for i := range k {
			n += remainders[i]
		}
	}

	return n
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	input := strings.Split(strings.TrimSpace(line), " ")
	n, _ := strconv.Atoi(input[0])
	k, _ := strconv.Atoi(input[1])

	writer.WriteString(strconv.Itoa(whatSum(n, k)))
	writer.WriteByte('\n')
}
