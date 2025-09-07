package main

import (
	"fmt"
	"strings"
)

func longestSubstring(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	alphabet := make(map[rune]int)
	for _, val := range s {
		alphabet[val]++
	}

	var skipLetter rune
	for key, val := range alphabet {
		if val < k {
			skipLetter = key
			break
		}
	}

	if skipLetter == 0 {
		return len(s)
	}

	var res int
	for _, val := range strings.Split(s, string(skipLetter)) {
		res = max(res, longestSubstring(val, k))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "aaabb"
	k := 3
	fmt.Println(longestSubstring(s, k))
}
