package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if !normalSymbol(s[l]) {
			l++
			continue
		}

		if !normalSymbol(s[r]) {
			r--
			continue
		}

		if getSymbol(s[l]) != getSymbol(s[r]) {
			return false
		}

		l++
		r--
	}

	return true
}

func normalSymbol(symbol byte) bool {
	return (symbol >= 'a' && symbol <= 'z') ||
		(symbol >= 'A' && symbol <= 'Z') ||
		(symbol >= '0' && symbol <= '9')
}

func getSymbol(symbol byte) byte {
	if symbol >= 'A' && symbol <= 'Z' {
		return symbol + ('Z' - 'A') + ('a' - 'Z')
	}
	return symbol
}

func main() {
	s := "1b1"
	fmt.Println(isPalindrome(s))
}
