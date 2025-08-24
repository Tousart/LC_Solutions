package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1) + 1
	n := len(word2) + 1
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = i
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for col := 1; col < n; col++ {
		for row := 1; row < m; row++ {
			if word1[row-1] == word2[col-1] {
				dp[row][col] = dp[row-1][col-1]
			} else {
				dp[row][col] = min(dp[row-1][col-1], min(dp[row][col-1], dp[row-1][col])) + 1
			}
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
