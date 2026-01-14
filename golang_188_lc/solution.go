package main

import "fmt"

func maxProfit(k int, prices []int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = []int{1001, 0}
	}

	for _, price := range prices {
		for i := 1; i < len(dp); i++ {
			dp[i][0] = min(dp[i][0], price-dp[i-1][1])
			dp[i][1] = max(dp[i][1], price-dp[i][0])
		}
	}

	return dp[k][1]
}

func main() {
	k := 2
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(k, prices))
}
