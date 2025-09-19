package main

import (
	"fmt"
	"strconv"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}

	for i := 1; i < len(num)-1; i++ {
		val := num[:i]
		if len(val) > 1 && val[0] == '0' {
			break
		}

		a, _ := strconv.Atoi(num[:i])

		for j := i + 1; j < len(num); j++ {
			val := num[i:j]
			if len(val) > 1 && val[0] == '0' {
				break
			}

			b, _ := strconv.Atoi(val)

			if move(a, b, num[j:]) {
				return true
			}
		}
	}

	return false
}

func move(a, b int, num string) bool {
	if len(num) == 0 {
		return true
	}

	sum := a + b

	for i := 1; i <= len(num); i++ {
		val := num[:i]
		if len(val) > 1 && val[0] == '0' {
			return false
		}

		digit, _ := strconv.Atoi(num[:i])

		if sum == digit {
			return move(b, digit, num[i:])
		} else if sum < digit {
			return false
		}
	}

	return false
}

func main() {
	num := "112358"
	fmt.Println(isAdditiveNumber(num))

	num1 := "199100199"
	fmt.Println(isAdditiveNumber(num1))

	num2 := "101"
	fmt.Println(isAdditiveNumber(num2))

	num3 := "0000"
	fmt.Println(isAdditiveNumber(num3))
}
