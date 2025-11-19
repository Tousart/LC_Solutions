package main

import "fmt"

func largestRectangleArea(heights []int) int {
	var res int
	stack := make([]int, 0)

	for i := range len(heights) + 1 {
		currHeigh := 0
		if i < len(heights) {
			currHeigh = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] > currHeigh {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			res = max(res, height*width)
		}

		stack = append(stack, i)
	}

	return res
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}
