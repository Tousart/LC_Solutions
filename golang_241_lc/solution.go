package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func diffWaysToCompute(expression string) []int {
	res := []int{}

	for i := 1; i < len(expression)-1; i++ {
		if !unicode.IsDigit(rune(expression[i])) {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, val1 := range left {
				for _, val2 := range right {
					switch expression[i] {
					case '+':
						res = append(res, val1+val2)
					case '-':
						res = append(res, val1-val2)
					case '*':
						res = append(res, val1*val2)
					default:
						res = append(res, val1/val2)
					}
				}
			}
		}
	}

	if len(res) == 0 {
		num, _ := strconv.Atoi(expression)
		res = append(res, num)
	}

	return res
}

func main() {
	expression := "2*3-4*5"
	fmt.Println(diffWaysToCompute(expression))
}
