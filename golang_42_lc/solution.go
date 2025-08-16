package main

import "fmt"

func trap(height []int) int {
	i, j := 1, len(height)-2
	left, right := height[0], height[len(height)-1]
	res := 0
	for i <= j {
		if height[i] > left {
			left = height[i]
			i++
			continue
		}

		if height[j] > right {
			right = height[j]
			j--
			continue
		}

		if left == right {
			if i != j {
				res += (left - height[i]) + (right - height[j])
			} else {
				res += left - height[i]
			}
			i++
			j--
		} else if left < right {
			res += left - height[i]
			i++
		} else {
			res += right - height[j]
			j--
		}
	}

	return res
}

func main() {
	height := []int{2, 0, 2}
	fmt.Println(trap(height))
}
