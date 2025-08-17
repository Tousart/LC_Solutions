package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
