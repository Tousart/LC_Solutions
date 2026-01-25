package main

import "fmt"

func minBitwiseArray(nums []int) []int {
	res := make([]int, len(nums))

	for i, num := range nums {
		j := 0
		for j < num {
			if (j | (j + 1)) == num {
				break
			}
			j++
		}

		if j == num {
			res[i] = -1
		} else {
			res[i] = j
		}
	}

	return res
}

func main() {
	nums := []int{2, 3, 5, 7}
	fmt.Println(minBitwiseArray(nums))
}
