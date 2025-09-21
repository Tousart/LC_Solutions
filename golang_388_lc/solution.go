package main

import (
	"fmt"
	"strings"
)

func lengthLongestPath(input string) int {
	var (
		res     int
		prev    int
		currLen int
	)
	stack := []string{}
	sliceInput := strings.Split(input, "\n")

	for _, name := range sliceInput {
		tabCount := strings.Count(name, "\t")
		if tabCount <= prev {
			for len(stack) > 0 {
				prevTabCount := strings.Count(stack[len(stack)-1], "\t")
				if tabCount > prevTabCount {
					break
				}
				prevElem := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				currLen -= len(prevElem) - prevTabCount
			}
		}

		stack = append(stack, name)
		currLen += len(name) - tabCount

		if strings.ContainsRune(name, '.') {
			res = max(res, currLen+len(stack)-1)
		}

		prev = tabCount
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	fmt.Println(lengthLongestPath(input))

	input1 := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	fmt.Println(lengthLongestPath(input1))
}
