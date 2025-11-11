package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func eblya(n int, reader *bufio.Reader, writer *bufio.Writer) {
	parents := make(map[string][]string)
	childs := make(map[string]string)

	for range n {
		line, _ := reader.ReadString('\n')
		input := strings.TrimSpace(line)

		if strings.Contains(input, " new ") {

			parent := input[5 : strings.Index(input, "=")-1]
			parents[parent] = strings.Split(input[strings.Index(input, "(")+1:len(input)-1], ",")
			childs[parent] = parent + ",0"

		} else if strings.Contains(input, "List ") {

			equallyIndex := strings.Index(input, "=")
			child := input[5 : equallyIndex-1]
			otchim := input[equallyIndex+2 : strings.Index(input, ".")]
			from, _ := strconv.Atoi(strings.Split(input[strings.Index(input, "(")+1:len(input)-1], ",")[0])

			commaIndex := strings.Index(childs[otchim], ",")
			parent := childs[otchim][:commaIndex]
			idx, _ := strconv.Atoi(childs[otchim][commaIndex+1:])

			childs[child] = parent + "," + strconv.Itoa(idx+from-1)

		} else if strings.Contains(input, ".set(") {

			pointIndex := strings.Index(input, ".")
			child := input[:pointIndex]
			ix := strings.Split(input[pointIndex+5:len(input)-1], ",")
			i, _ := strconv.Atoi(ix[0])
			x := ix[1]

			commaIndex := strings.Index(childs[child], ",")
			parent := childs[child][:commaIndex]
			idx, _ := strconv.Atoi(childs[child][commaIndex+1:])

			parents[parent][idx+i-1] = x

		} else if strings.Contains(input, ".add(") {

			pointIndex := strings.Index(input, ".")
			parent := input[:pointIndex]
			x := input[pointIndex+5 : len(input)-1]

			parents[parent] = append(parents[parent], x)

		} else if strings.Contains(input, ".get(") {

			pointIndex := strings.Index(input, ".")
			child := input[:pointIndex]
			iString := input[pointIndex+5 : len(input)-1]
			i, _ := strconv.Atoi(iString)

			commaIndex := strings.Index(childs[child], ",")
			parent := childs[child][:commaIndex]
			idx, _ := strconv.Atoi(childs[child][commaIndex+1:])

			writer.WriteString(parents[parent][idx+i-1])
			writer.WriteByte('\n')
			writer.Flush()

		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	eblya(n, reader, writer)
}
