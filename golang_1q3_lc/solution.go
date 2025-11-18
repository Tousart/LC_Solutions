package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	l, r, zeroes := 0, 0, 0
	for r < len(nums) {
		if nums[r] == 0 {
			zeroes++
		}

		if nums[l] == 1 {
			if zeroes > 0 {
				l++
			}
		} else {
			zeroes--
			l++
		}

		r++
	}
	return r - l
}

func main() {
	nums := []int{1, 0, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
