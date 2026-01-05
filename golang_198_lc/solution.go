package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	nums[2] = max(nums[2]+nums[0], nums[1])

	for i := 3; i < len(nums); i++ {
		nums[i] = max(max(nums[i]+nums[i-2], nums[i-1]), nums[i]+nums[i-3])
	}

	return nums[len(nums)-1]
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))

	nums1 := []int{100, 1, 1, 100}
	fmt.Println(rob(nums1))
}
