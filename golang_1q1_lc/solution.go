package main

import "fmt"

func getConcatenation(nums []int) []int {
	res := make([]int, len(nums)*2)
	i := 0
	for i < len(nums) {
		res[i] = nums[i]
		i++
	}
	for i < len(nums)*2 {
		res[i] = nums[i-len(nums)]
		i++
	}
	return res
}

func main() {
	nums := []int{1, 2, 4}
	fmt.Println(getConcatenation(nums))
}
