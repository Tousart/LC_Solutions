package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	var res int
	sort.Ints(nums)
	for i := 2; i < len(nums); i++ {
		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				res += r - l
				r--
			} else {
				l++
				// left, right := l, r
				// for left < right {
				// 	mid := (left + right) / 2
				// 	if nums[mid]+nums[r] > nums[i] {
				// 		right = mid
				// 	} else {
				// 		left = mid + 1
				// 	}
				// }
				// l = left
			}
		}
	}
	return res
}

func main() {
	nums := []int{4, 2, 3, 4}
	fmt.Println(triangleNumber(nums))
}
