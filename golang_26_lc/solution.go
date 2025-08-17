package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	i, j := 0, 1
	for j < len(nums) {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
		j++
	}

	return i + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(nums, removeDuplicates(nums))
}
