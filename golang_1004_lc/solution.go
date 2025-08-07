package main

import "fmt"

func longestOnes(nums []int, k int) int {
	var (
		i int
		j int
	)

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

	return j - i
}

// func longestOnes(nums []int, k int) int {
// 	var (
// 		ans     int
// 		zeroInd int
// 	)

// 	zeros := []int{-1}

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			zeros = append(zeros, i)
// 			if k == 0 {
// 				ans = maximum(ans, i-zeros[zeroInd]-1)
// 				zeroInd++
// 			} else {
// 				k--
// 			}
// 		}
// 	}

// 	ans = maximum(ans, len(nums)-zeros[zeroInd]-1)

// 	return ans
// }

// func maximum(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k := 0
	fmt.Println(longestOnes(nums, k))
	nums1 := []int{0, 0, 1, 1, 1, 0, 0}
	k1 := 0
	fmt.Println(longestOnes(nums1, k1))
}
