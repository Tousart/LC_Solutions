package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	nums1 := strings.Split(version1, ".")
	nums2 := strings.Split(version2, ".")
	n := max(len(nums1), len(nums2))
	for i := range n {
		num1 := 0
		if i < len(nums1) {
			num1, _ = strconv.Atoi(nums1[i])
		}
		num2 := 0
		if i < len(nums2) {
			num2, _ = strconv.Atoi(nums2[i])
		}
		if num1 > num2 {
			return 1
		} else if num2 > num1 {
			return -1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	version1 := "1.2"
	version2 := "1.10"
	fmt.Println(compareVersion(version1, version2))
}
