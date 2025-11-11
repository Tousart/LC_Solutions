package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func kOutput(k int, tasks []string) string {
	res := []string{}
	themes := make(map[string]int)
	for _, task := range tasks {
		if count := themes[task]; count == 0 {
			res = append(res, task)
		}

		if len(res) == k {
			return strings.Join(res, " ")
		}

		themes[task]++
	}

	for task := range themes {
		for themes[task] > 1 {
			res = append(res, task)
			themes[task]--

			if len(res) == k {
				return strings.Join(res, " ")
			}
		}
	}

	return strings.Join(res, " ")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.Split(strings.TrimSpace(line1), " ")[1])

	line2, _ := reader.ReadString('\n')
	tasks := strings.Split(strings.TrimSpace(line2), " ")

	result := kOutput(k, tasks)
	writer.WriteString(result)
	writer.WriteByte('\n')
}
