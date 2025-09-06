package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{})
	maxLength := len(wordDict[0])
	for _, val := range wordDict {
		maxLength = max(maxLength, len(val))
		dict[val] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := range len(s) {
		for j := i; j < len(s); j++ {
			if (j - i + 1) > maxLength {
				break
			}

			if _, ok := dict[s[i:j+1]]; ok && !dp[j+1] {
				dp[j+1] = dp[i]
			}
		}
	}

	return dp[len(s)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "applepenapple"
	wordDict := []string{
		"apple",
		"pen",
	}
	fmt.Println(wordBreak(s, wordDict))

	s1 := "aaaaaaa"
	wordDict1 := []string{
		"aaa",
		"aaaa",
	}
	fmt.Println(wordBreak(s1, wordDict1))
}
