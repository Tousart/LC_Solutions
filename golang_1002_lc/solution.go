package main

import (
	"fmt"
	"strings"
)

func commonChars(words []string) []string {
	letters := make([]int, 26)

	for _, letter := range words[0] {
		letters[letter-'a']++
	}

	for i := 1; i < len(words); i++ {
		for j := range 26 {
			letters[j] = min(letters[j], strings.Count(words[i], string(rune(j+'a'))))
		}
	}

	res := make([]string, 0)
	for i, letterCount := range letters {
		for range letterCount {
			res = append(res, string(rune(i+'a')))
		}
	}

	return res
}

func main() {
	words := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(words))
}
