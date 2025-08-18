package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	prefix := 0
	remains := map[int]int{0: -1}

	for i, val := range nums {
		prefix += val
		if ind, ok := remains[prefix%k]; ok {
			if i-ind >= 2 {
				return true
			}
		} else {
			remains[prefix%k] = i
		}
	}
	return false
}

func main() {
	nums := []int{23, 2, 4, 6, 6}
	fmt.Println(checkSubarraySum(nums, 7))

	nums1 := []int{0, 0}
	fmt.Println(checkSubarraySum(nums1, 1))
}
