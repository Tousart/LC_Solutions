package main

import "fmt"

func subarraySum(nums []int, k int) int {
	var (
		count int
		sum   int
	)
	sums := map[int]int{0: 1}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		count += sums[sum-k]
		sums[sum] += 1
	}

	return count
}

func main() {
	nums := []int{1, -1, 0}
	k := 0
	fmt.Println(subarraySum(nums, k))
}
