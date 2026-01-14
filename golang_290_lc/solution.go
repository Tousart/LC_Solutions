package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	seq := strings.Split(s, " ")

	if len(pattern) != len(seq) {
		return false
	}

	wordLet := make(map[string]byte)
	letExists := make(map[byte]struct{})

	for i, word := range seq {
		if let, ok := wordLet[word]; ok && let != pattern[i] {
			return false
		} else if !ok {
			if _, ok := letExists[pattern[i]]; ok {
				return false
			}

			wordLet[word] = pattern[i]
			letExists[pattern[i]] = struct{}{}
		}
	}

	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat d"
	fmt.Println(wordPattern(pattern, s))
}
