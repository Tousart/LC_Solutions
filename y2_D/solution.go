package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func makeSentence(s string, words map[string]struct{}) string {
	childParent := make(map[int]int)
	parents := []int{0}

	for i := 1; i <= len(s); i++ {
		for _, parent := range parents {
			if _, ok := words[s[parent:i]]; ok {
				childParent[i] = parent
				parents = append(parents, i)
				break
			}
		}
	}

	res := []string{}
	current := len(s)
	for current != 0 {
		parent := childParent[current]
		res = append(res, s[parent:current])
		current = parent
	}

	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return strings.Join(res, " ")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	s := strings.TrimSpace(line1)

	line2, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line2))

	words := make(map[string]struct{})
	for range n {
		line, _ := reader.ReadString('\n')
		words[strings.TrimSpace(line)] = struct{}{}
	}

	writer.WriteString(makeSentence(s, words))
}
