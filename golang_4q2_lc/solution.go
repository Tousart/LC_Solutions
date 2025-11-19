package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))

	for i, temperature := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperature {
			idx := stack[len(stack)-1]
			res[idx] = i - idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	temperatures := []int{30, 40, 50, 60}
	fmt.Println(dailyTemperatures(temperatures))
}
