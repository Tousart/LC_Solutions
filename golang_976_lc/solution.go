package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-1]+nums[i-2] > nums[i] {
			return nums[i-1] + nums[i-2] + nums[i]
		}
	}
	return 0
}

func main() {
	nums := []int{2, 1, 2}
	fmt.Println(largestPerimeter(nums))
}
