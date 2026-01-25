package main

import "fmt"

func minBitwiseArray(nums []int) []int {
	res := make([]int, len(nums))

	for i, num := range nums {
		one := 1
		val := -1

		for (num & one) != 0 {
			val = num - one
			one <<= 1
		}

		res[i] = val
	}

	return res
}

func main() {
	nums := []int{2, 3, 5, 7}
	fmt.Println(minBitwiseArray(nums))
}

// func minBitwiseArray(nums []int) []int {
// 	res := make([]int, len(nums))

// 	for i, num := range nums {
// 		if (num & 1) == 0 {
// 			res[i] = -1
// 			continue
// 		}

// 		for j := 1; j < 32; j++ {
// 			if (num & (1 << j)) == 0 {
// 				res[i] = num ^ (1 << (j - 1))
// 				break
// 			}
// 		}
// 	}

// 	return res
// }
