package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	letIdxs := make(map[rune]int)
	idx, res := 0, 0

	for i, letter := range s {
		if letIdx, ok := letIdxs[letter]; ok && idx <= letIdx {
			res = max(res, i-idx)
			idx = letIdx + 1
		}
		letIdxs[letter] = i
	}

	return max(res, len(s)-idx)
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
