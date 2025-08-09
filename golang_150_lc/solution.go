package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, val := range tokens {
		if val == "+" || val == "-" || val == "*" || val == "/" {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			if val == "+" {
				stack[len(stack)-2] = a + b
			} else if val == "-" {
				stack[len(stack)-2] = a - b
			} else if val == "*" {
				stack[len(stack)-2] = a * b
			} else {
				stack[len(stack)-2] = a / b
			}
			stack = stack[:len(stack)-1]
		} else {
			intVal, _ := strconv.Atoi(val)
			stack = append(stack, intVal)
		}
	}
	return stack[0]
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
}
