package main

import "fmt"

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([]int, len(s1)+1)
	for i := 1; i <= len(s1); i++ {
		dp[i] = dp[i-1] + int(s1[i-1])
	}

	newDp := make([]int, len(s1)+1)
	for row := range len(s2) {
		newDp[0] = dp[0] + int(s2[row])

		for col := 1; col <= len(s1); col++ {
			if s2[row] == s1[col-1] {
				newDp[col] = dp[col-1]
			} else {
				newDp[col] = min(dp[col]+int(s2[row]), newDp[col-1]+int(s1[col-1]))
			}
		}

		dp, newDp = newDp, dp
	}

	return dp[len(s1)]
}

func main() {
	s1 := "delete"
	s2 := "leet"
	fmt.Println(minimumDeleteSum(s1, s2))
}
