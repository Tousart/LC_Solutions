package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	inds := make(map[rune]int)
	j, length := 0, 0

	for i, val := range s {
		if ind, ok := inds[val]; ok && ind >= j {
			length = maximum(length, i-j)
			j = ind + 1
		}
		inds[val] = i
	}

	return maximum(length, len(s)-j)
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "abba"
	fmt.Println(lengthOfLongestSubstring(s))
}
