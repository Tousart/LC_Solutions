package main

import "fmt"

func minimumPairRemoval(nums []int) int {
	var res int

	for {
		idx := 0
		minSum := 50001
		stop := true

		for i := range len(nums) - 1 {
			if (nums[i] + nums[i+1]) < minSum {
				idx = i
				minSum = nums[i] + nums[i+1]
			}

			if nums[i+1] < nums[i] {
				stop = false
			}
		}

		if stop {
			break
		}

		nums[idx] = minSum
		temp := nums[idx+2:]
		nums = append(nums[:idx+1], temp...)
		res++
	}

	return res
}

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(minimumPairRemoval(nums))

	nums1 := []int{689, -360, 234, 673, 663, -741, 480, 860, -707, 209, 246, 792, 930, 696, -305}
	fmt.Println(minimumPairRemoval(nums1))
}
