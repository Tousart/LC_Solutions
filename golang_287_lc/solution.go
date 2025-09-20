package main

import "fmt"

func findDuplicate(nums []int) int {
	slow := nums[nums[0]]
	fast := nums[nums[nums[0]]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

func findDuplicate1(nums []int) int {
	for _, num := range nums {
		ind := num
		if num < 0 {
			ind *= -1
		}

		if nums[ind] < 0 {
			return ind
		}

		nums[ind] *= -1
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate1(nums))

	nums1 := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums1))
}
