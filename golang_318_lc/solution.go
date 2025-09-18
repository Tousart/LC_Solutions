package main

import (
	"fmt"
	// "strings"
)

func maxProduct(words []string) int {
	var res int
	masks := make([]uint32, len(words))
	for i, word := range words {
		masks[i] = makeBitMask(word)
		for j := range i {
			if masks[i]&masks[j] == 0 {
				res = max(res, len(words[j])*len(word))
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func makeBitMask(word string) uint32 {
	var mask uint32
	// for i := 'a'; i <= 'z'; i++ {
	// 	if strings.ContainsRune(word, i) {
	// 		mask |= 1 << (i - 'a')
	// 	}
	// }

	for _, letter := range word {
		mask |= 1 << (letter - 'a')
	}

	return mask
}

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(maxProduct(words))
}
