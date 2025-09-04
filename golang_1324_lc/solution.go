package main

import (
	"fmt"
	"strings"
)

func printVertically(s string) []string {
	res := []string{}
	sliceS := strings.Split(s, " ")
	n := 0
	for _, val := range sliceS {
		n = max(n, len(val))
	}

	for i := range n {
		word := []byte{}
		for _, val := range sliceS {
			elem := byte(' ')
			if i < len(val) {
				elem = val[i]
			}
			word = append(word, elem)
		}
		res = append(res, strings.TrimRight(string(word), " "))
	}

	fmt.Println(len(res))
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "CONTEST IS COMING"
	fmt.Println(printVertically(s))
}
