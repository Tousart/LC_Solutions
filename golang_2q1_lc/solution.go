package main

import "fmt"

func findErrorNums(nums []int) []int {
	numsMap := make(map[int]struct{})
	expected, actual := len(nums)*(len(nums)+1)/2, 0
	badNum := 0
	for _, num := range nums {
		actual += num
		if _, ok := numsMap[num]; ok {
			badNum = num
		}
		numsMap[num] = struct{}{}
	}
	return []int{badNum, expected - actual + badNum}
}

func main() {
	nums := []int{1, 2, 2, 4}
	fmt.Println(findErrorNums(nums))
}
