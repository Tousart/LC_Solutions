package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)

	res := 0

	i, j := 0, len(nums)-1
	for i < j {
		res = max(res, nums[i]+nums[j])
		i++
		j--
	}

	return res
}

func main() {
	nums := []int{3, 5, 2, 3}
	fmt.Println(minPairSum(nums))
}
