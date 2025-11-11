package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func makeString(s string, pieces map[string][]string, step, size int) string {
	res := make([]string, size)
	idx := 0
	for i := 0; i < len(s); i += step {
		key := s[i : i+step]
		res[idx] = pieces[key][0]
		pieces[key] = pieces[key][1:]
		idx++
	}
	return strings.Join(res, " ")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	nm := strings.Split(strings.TrimSpace(line1), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	line2, _ := reader.ReadString('\n')
	s := strings.TrimSpace(line2)

	pieces := make(map[string][]string)
	for i := range m {
		line, _ := reader.ReadString('\n')
		key := strings.TrimSpace(line)
		if pieces[key] == nil {
			pieces[key] = []string{}
		}
		pieces[key] = append(pieces[key], strconv.Itoa(i+1))
	}

	writer.WriteString(makeString(s, pieces, n/m, m))
}
