package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	right := len(s) - 1
	for s[right] == ' ' {
		right--
	}

	left := right - 1
	for left >= 0 && s[left] != ' ' {
		left--
	}

	return right - left
}

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))

	s1 := "a"
	fmt.Println(lengthOfLastWord(s1))
}
