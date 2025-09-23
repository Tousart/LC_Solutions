package main

import (
	"fmt"
	"sort"
)

func longestWord(words []string) string {
	var res string

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) < len(words[j])
	})

	canBuild := make(map[string]struct{})
	for _, word := range words {
		if _, ok := canBuild[word[:len(word)-1]]; ok || len(word) == 1 {
			canBuild[word] = struct{}{}
			if (len(word) > len(res)) || (len(word) == len(res) && word < res) {
				res = word
			}
		}
	}

	return res
}

func main() {
	words := []string{"w", "wo", "wor", "worl", "world"}
	fmt.Println(longestWord(words))

	words1 := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println(longestWord(words1))
}
