package main

import (
	"fmt"
	"sort"
)

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	dpMap := make(map[string]int)
	for _, word := range words {
		for idx := range word {
			dpMap[word] = max(dpMap[word], 1+dpMap[word[:idx]+word[idx+1:]])
		}
	}

	var res int
	for _, val := range dpMap {
		res = max(res, val)
	}

	return res
}

func main() {
	words := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	fmt.Println(longestStrChain(words))
}
