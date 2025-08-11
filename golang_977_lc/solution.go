package main

import "fmt"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j, arrInd := 0, len(nums)-1, len(nums)-1

	for i <= j {
		a := nums[i] * nums[i]
		b := nums[j] * nums[j]
		if a > b {
			res[arrInd] = a
			i++
		} else {
			res[arrInd] = b
			j--
		}
		arrInd--
	}
	return res
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
