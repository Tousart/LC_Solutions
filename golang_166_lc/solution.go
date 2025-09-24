package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	res := []byte{}
	if (numerator < 0) != (denominator < 0) {
		res = append(res, '-')
	}

	if numerator < 0 {
		numerator *= -1
	}

	if denominator < 0 {
		denominator *= -1
	}

	res = append(res, []byte(strconv.Itoa(numerator/denominator))...)

	remainder := numerator % denominator
	if remainder == 0 {
		return string(res)
	}

	res = append(res, '.')

	remainderMap := make(map[int]int)
	for remainder != 0 {
		fmt.Println(string(res))
		if idx, ok := remainderMap[remainder]; ok {
			res = append(res[:idx], []byte("("+string(res[idx:])+")")...)
			return string(res)
		}

		remainderMap[remainder] = len(res)
		remainder *= 10
		res = append(res, byte('0'+remainder/denominator))
		remainder %= denominator
	}

	return string(res)
}

func main() {
	numerator := 4
	denomerator := 333
	fmt.Println(fractionToDecimal(numerator, denomerator))

	numerator1 := 5
	denomerator1 := 2
	fmt.Println(fractionToDecimal(numerator1, denomerator1))

	numerator2 := 4
	denomerator2 := 2
	fmt.Println(fractionToDecimal(numerator2, denomerator2))

	numerator3 := 1
	denomerator3 := 6
	fmt.Println(fractionToDecimal(numerator3, denomerator3))

	numerator4 := 1
	denomerator4 := -1
	fmt.Println(fractionToDecimal(numerator4, denomerator4))
}
