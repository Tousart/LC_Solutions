package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	inds := make(map[int]int)

	for i, val := range nums {
		diff := target - val
		if ind, ok := inds[diff]; ok {
			return []int{ind, i}
		}
		inds[val] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

// func twoSum(nums []int, target int) []int {
// 	temp := make([]int, len(nums))
// 	copy(temp, nums)
// 	res := []int{0, 0}

// 	sort.Ints(temp)

// 	i, j := 0, len(temp)-1
// 	for i < len(temp) && j >= 0 {
// 		sum := temp[i] + temp[j]
// 		if sum == target {
// 			res[0] = temp[i]
// 			res[1] = temp[j]
// 			break
// 		} else if sum < target {
// 			i++
// 		} else {
// 			j--
// 		}
// 	}

// 	for k := range nums {
// 		if res[0] == nums[k] {
// 			res[0] = k
// 			break
// 		}
// 	}

// 	for k := len(nums) - 1; k >= 0; k-- {
// 		if res[1] == nums[k] {
// 			res[1] = k
// 			break
// 		}
// 	}

// 	return res
// }
