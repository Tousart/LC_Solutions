package main

import "fmt"

func maxFrequencyElements(nums []int) int {
	seq := make(map[int]int)
	maxFreq := 0
	for _, num := range nums {
		seq[num]++
		maxFreq = max(maxFreq, seq[num])
	}

	res := 0
	for _, val := range seq {
		if val == maxFreq {
			res += maxFreq
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 2, 3, 1, 4}
	fmt.Println(maxFrequencyElements(nums))
}
