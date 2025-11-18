package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	freq := make([]int, 101)
	for _, num := range nums {
		freq[num]++
	}
	for i, num := range nums {
		sum := 0
		for j := range num {
			sum += freq[j]
		}
		nums[i] = sum
	}
	return nums
}

func main() {
	nums := []int{6, 5, 4, 8}
	fmt.Println(smallerNumbersThanCurrent(nums))
}
