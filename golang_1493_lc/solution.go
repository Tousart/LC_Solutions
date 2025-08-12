package main

import "fmt"

func longestSubarray(nums []int) int {
	i, j, k := 0, 0, 1

	for j < len(nums) {
		if nums[j] == 0 {
			k--
		}
		j++

		if k < 0 {
			if nums[i] == 0 {
				k++
			}
			i++
		}
	}

	return j - i - 1
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}
