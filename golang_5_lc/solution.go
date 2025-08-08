package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	var (
		maxInd    int
		maxLength int
		r         int
		l         int
	)
	newS := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(newS)
	pLengths := make([]int, n)

	for i := 1; i < n-1; i++ {
		if i < r {
			pLengths[i] = minimum(pLengths[l+r-i], r-i)
		}

		for newS[i-pLengths[i]-1] == newS[i+pLengths[i]+1] {
			pLengths[i]++
		}

		if i+pLengths[i] > r {
			r = i + pLengths[i]
			l = i - pLengths[i]
		}
	}

	for i, val := range pLengths {
		if val > maxLength {
			maxLength = val
			maxInd = i
		}
	}

	return s[(maxInd-maxLength)/2 : (maxInd+maxLength)/2]
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
