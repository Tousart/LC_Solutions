package main

import "fmt"

func specialTriplets(nums []int) int {
	var res int

	mod := 1000000007

	x2 := make(map[int]int)
	counts := make([]int, len(nums))

	for i, num := range nums {
		counts[i] = x2[num*2]
		x2[num]++
	}

	for i := 1; i < len(nums)-1; i++ {
		k := nums[i] * 2
		val := x2[k]
		if k == 0 {
			val--
		}
		res += counts[i] * (val - counts[i])
	}

	return res % mod
}

func main() {
	nums := []int{8, 4, 2, 8, 4}
	fmt.Println(specialTriplets(nums))

	nums1 := []int{6, 3, 6}
	fmt.Println(specialTriplets(nums1))

	nums2 := []int{0, 1, 0, 0}
	fmt.Println(specialTriplets(nums2))

	nums3 := []int{0, 0, 0, 0}
	fmt.Println(specialTriplets(nums3))
}
