package main

import "fmt"

func plusOne(digits []int) []int {
	res := make([]int, len(digits)+1)
	res[0] = 1
	remain := 1

	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + remain
		res[i+1] = sum % 10
		remain = sum / 10
	}

	if remain == 1 {
		return res
	}

	return res[1:]
}

func main() {
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}
