package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var (
		ans   int
		start int = 1
	)
	mapS := make(map[rune]int)

	for ind, val := range s {
		oldInd := mapS[val]
		if oldInd != 0 && oldInd >= start {
			ans = maximum(ans, ind+1-start)
			start = oldInd + 1
		}
		mapS[val] = ind + 1
	}

	ans = maximum(ans, len(s)+1-start)

	return ans
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
