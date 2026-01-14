package main

import "fmt"

func maxDotProduct(nums1 []int, nums2 []int) int {
	m, n := len(nums1)+1, len(nums2)+1
	dp := make([]int, n)
	prev := make([]int, n)
	for col := range n {
		dp[col], prev[col] = -1000001, -1000001
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[col] = max(
				nums1[row-1]*nums2[col-1],
				nums1[row-1]*nums2[col-1]+max(prev[col-1], 0),
				prev[col],
				dp[col-1],
			)
		}

		dp, prev = prev, dp
	}

	return prev[n-1]
}

func main() {
	nums1 := []int{-1, -1}
	nums2 := []int{1, 1}
	fmt.Println(maxDotProduct(nums1, nums2))
}
