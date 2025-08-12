package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var (
		i      int
		j      int
		strVal string
		res    []string
	)

	for j < len(nums) {
		if j+1 == len(nums) || nums[j+1]-nums[j] > 1 {
			if i == j {
				strVal = strconv.Itoa(nums[i])
			} else {
				strVal = strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j])
			}
			res = append(res, strVal)
			i = j + 1
		}
		j++
	}

	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
