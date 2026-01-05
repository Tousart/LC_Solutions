package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)

	for _, word := range words {
		letPat := make(map[byte]byte)
		patLet := make(map[byte]byte)
		idx := 0

		for idx < len(word) {
			if pat, ok := letPat[word[idx]]; !ok {
				if _, ok := patLet[pattern[idx]]; !ok {
					letPat[word[idx]] = pattern[idx]
					patLet[pattern[idx]] = word[idx]
				} else {
					break
				}
			} else {
				if pat != pattern[idx] {
					break
				}
			}

			idx++
		}

		if idx == len(word) {
			res = append(res, word)
		}
	}

	return res
}

func main() {
	words := []string{"abc", "aqq", "ccc"}
	pattern := "mee"
	fmt.Println(findAndReplacePattern(words, pattern))
}
