package main

import (
	"fmt"
	"math"
)

func maxSubarraySum(nums []int, k int) int64 {
	minPrefix := make([]int, k)

	for i := 1; i < k; i++ {
		minPrefix[i] = math.MaxInt64
	}

	var (
		res    int = math.MinInt64
		prefix int
	)

	for i, num := range nums {
		prefix += num
		idx := (i + 1) % k

		if minPrefix[idx] != math.MaxInt64 {
			res = max(res, prefix-minPrefix[idx])
		}

		if prefix < minPrefix[idx] {
			minPrefix[idx] = prefix
		}
	}

	return int64(res)
}

func main() {
	nums := []int{-5, 1, 2, -3, 4}
	k := 2
	fmt.Println(maxSubarraySum(nums, k))
}
