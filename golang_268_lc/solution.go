package main

import "fmt"

func missingNumber(nums []int) int {
	sum1, sum2 := 0, 0
	for i := 0; i < len(nums); i++ {
		sum1 += nums[i]
		sum2 += i
	}

	return sum2 + len(nums) - sum1
}

func main() {
	nums := []int{0, 1, 3}
	fmt.Println(missingNumber(nums))
}
