package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	iElem := nums[0]
	jElem := math.MaxInt64
	for i := 1; i < len(nums); i++ {
		if nums[i] < iElem {
			iElem = nums[i]
			continue
		} else if nums[i] > iElem && nums[i] < jElem {
			jElem = nums[i]
			continue
		} else if nums[i] > jElem {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{100, 5, 101}
	fmt.Println(increasingTriplet(nums))
}
