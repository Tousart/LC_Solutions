package main

import "fmt"

func repeatedNTimes(nums []int) int {
	freqs := make([]int, 10001)
	for _, num := range nums {
		if freqs[num] == 1 {
			return num
		}
		freqs[num]++
	}
	return 0
}

func main() {
	nums := []int{1, 2, 3, 3}
	fmt.Println(repeatedNTimes(nums))
}
