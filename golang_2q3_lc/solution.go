package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	res := []int{}
	for i := range nums {
		if nums[i]-1 != i {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
