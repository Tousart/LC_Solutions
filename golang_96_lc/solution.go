package main

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := range i {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

func main() {
	n := 3
	fmt.Println(numTrees(n))
}
