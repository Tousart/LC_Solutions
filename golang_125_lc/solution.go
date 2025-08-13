package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var (
		i int
		j int = len(s) - 1
	)

	runesS := []rune(s)

	for i <= j {
		for i < j {
			if unicode.IsLetter(runesS[i]) || unicode.IsDigit(runesS[i]) {
				break
			}
			i++
		}

		for j > i {
			if unicode.IsLetter(runesS[j]) || unicode.IsDigit(runesS[j]) {
				break
			}
			j--
		}

		if !strings.EqualFold(strings.ToLower(string(s[i])), strings.ToLower(string(s[j]))) {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
