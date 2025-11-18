package main

import "fmt"

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	i, j := 0, 0
	for i < n {
		res[j] = nums[i]
		res[j+1] = nums[i+n]
		j += 2
		i++
	}
	return res
}

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	fmt.Println(shuffle(nums, n))
}
