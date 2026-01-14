package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	currCombination := make([]int, 0)

	sort.Ints(nums)

	var makeCombination func(idx int)
	makeCombination = func(idx int) {
		comb := make([]int, len(currCombination))
		copy(comb, currCombination)
		res = append(res, comb)

		if idx == len(nums) {
			return
		}

		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}

			currCombination = append(currCombination, nums[i])
			makeCombination(i + 1)
			currCombination = currCombination[:len(currCombination)-1]
		}
	}

	makeCombination(0)

	return res
}

func main() {
	nums := []int{4, 4, 4, 1, 4}
	fmt.Println(subsetsWithDup(nums))
}
