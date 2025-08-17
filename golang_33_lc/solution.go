package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		midVal := nums[mid]
		if midVal == target {
			return mid
		}

		leftVal := nums[left]
		rightVal := nums[right]
		if leftVal < rightVal {
			if midVal > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if midVal >= leftVal {
			if midVal > target && target >= leftVal {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if midVal < target && target <= rightVal {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))

	nums1 := []int{3, 1}
	target1 := 1
	fmt.Println(search(nums1, target1))
}
