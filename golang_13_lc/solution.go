package main

import (
	"fmt"
)

func romanToInt(s string) int {
	letters := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	lastVal := 0
	for _, letter := range s {
		if lastVal >= letters[letter] {
			lastVal = letters[letter]
		} else {
			res -= lastVal
			lastVal = letters[letter] - lastVal
		}
		res += lastVal
	}

	return res
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
