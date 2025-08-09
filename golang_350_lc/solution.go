package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	numsMap := make(map[int]int)
	for _, val := range nums1 {
		numsMap[val]++
	}

	for _, val := range nums2 {
		if numsMap[val] != 0 {
			numsMap[val]--
			res = append(res, val)
		}
	}

	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersect(nums1, nums2))
}
