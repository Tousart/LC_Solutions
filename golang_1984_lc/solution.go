package main

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)

	res := 100001

	for i := range len(nums) - k + 1 {
		res = min(res, nums[i+k-1]-nums[i])
	}

	return res
}

func main() {
	nums := []int{9, 4, 1, 7}
	k := 2
	fmt.Println(minimumDifference(nums, k))
}
