package main

import (
	"fmt"
)

func peopleAwareOfSecret(n int, delay int, forget int) int {
	mod := int(1e9 + 7)
	dp := make([]int, n)
	dp[0] = 1

	for i := range n {
		for j := i + delay; j < min(n, i+forget); j++ {
			dp[j] += dp[i]
			dp[j] %= mod
		}
	}

	res := 0
	for i := n - forget; i < n; i++ {
		res += dp[i]
		res %= mod
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 6
	delay := 2
	forget := 4
	fmt.Println(peopleAwareOfSecret(n, delay, forget))
}
