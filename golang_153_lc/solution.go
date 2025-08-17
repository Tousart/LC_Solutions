package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right-1 {
		mid := (left + right) / 2
		if nums[left] < nums[right] {
			return nums[left]
		} else if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return minimum(nums[left], nums[right])
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{4, 5, 1, 2, 3}
	fmt.Println(findMin(nums))
}

// func findMin(nums []int) int {
// 	left, right := 0, len(nums)-1
// 	for left < right-1 {
// 		mid := (left + right) / 2
// 		midVal := nums[mid]

// 		leftVal := nums[left]
// 		rightVal := nums[right]
// 		if leftVal < rightVal {
// 			return leftVal
// 		} else if midVal > leftVal {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}

// 	return minimum(nums[left], nums[right])
// }
