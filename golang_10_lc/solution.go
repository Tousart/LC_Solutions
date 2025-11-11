package main

import "fmt"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true

	for col := 1; col < len(p)+1; col++ {
		if p[col-1] == '*' {
			dp[0][col] = dp[0][col-2]
		}
	}

	for row := 1; row < len(s)+1; row++ {
		for col := 1; col < len(p)+1; col++ {
			switch p[col-1] {
			case s[row-1], '.':
				dp[row][col] = dp[row-1][col-1]
			case '*':
				dp[row][col] = dp[row][col-2]
				if (p[col-2] == s[row-1] || p[col-2] == '.') && dp[row-1][col] {
					dp[row][col] = true
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}

func main() {
	s := "aa"
	p := ".*"
	fmt.Println(isMatch(s, p))
}
