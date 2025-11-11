package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findHappiness(n int, mushrooms []string) int {
	var (
		res  int
		minV int = 1001
		maxM int
	)

	for i := range n {
		val, _ := strconv.Atoi(mushrooms[i])
		if i%2 != 0 {
			res -= val
			maxM = max(maxM, val)
		} else {
			res += val
			minV = min(minV, val)
		}
	}

	if maxM > minV {
		return res + 2*(maxM-minV)
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
	writer.WriteString(strconv.Itoa(findHappiness(n, strings.Split(strings.TrimSpace(line2), " "))))
}
