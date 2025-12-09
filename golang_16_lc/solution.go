package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var res int

	sort.Ints(nums)

	minDiff := math.MaxInt64
	for i := range len(nums) - 2 {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			} else {
				diff := int(math.Abs(float64(sum - target)))
				if diff < minDiff {
					minDiff = diff
					res = sum
				}

				if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))

	nums1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	target1 := 1
	fmt.Println(threeSumClosest(nums1, target1))

	nums2 := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	target2 := -2
	fmt.Println(threeSumClosest(nums2, target2))

	nums3 := []int{2, 3, 8, 9, 10}
	target3 := 16
	fmt.Println(threeSumClosest(nums3, target3))
}
