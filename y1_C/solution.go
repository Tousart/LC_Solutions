package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func countOfWeeks(password string) int {
	letters := make(map[byte]int)
	for i := range password {
		letters[password[i]]++
	}

	res := 1

	for i := range password {
		letters[password[i]]--
		for key, val := range letters {
			if key != password[i] {
				res += val
			}
		}
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')

	writer.WriteString(strconv.Itoa(countOfWeeks(strings.TrimSpace(line))))
}
