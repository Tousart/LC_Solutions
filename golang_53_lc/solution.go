package main

import "fmt"

func maxSubArray(nums []int) int {
	curr := nums[0]
	res := curr

	for i := 1; i < len(nums); i++ {
		curr = max(nums[i], nums[i]+curr)
		res = max(res, curr)
	}

	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
