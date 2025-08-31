package main

import "fmt"

func plusOne(digits []int) []int {
	res := []int{0}
	res = append(res, digits...)
	shift := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + shift
		res[i+1] = sum % 10
		shift = sum / 10

		if shift == 0 {
			return res[1:]
		}
	}
	res[0] = 1
	return res
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))

	digits1 := []int{1, 9, 9, 9}
	fmt.Println(plusOne(digits1))
}
