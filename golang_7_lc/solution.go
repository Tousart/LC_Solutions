package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	xStr := ""
	if x < 0 {
		xStr = strconv.Itoa(x * -1)
	} else {
		xStr = strconv.Itoa(x)
	}
	bytes := []byte{}

	for i := len(xStr) - 1; i >= 0; i-- {
		if xStr[i] == '0' && len(bytes) == 0 {
			continue
		}

		bytes = append(bytes, xStr[i])

		if len(bytes) > 10 {
			return 0
		}
	}

	if len(bytes) == 10 {
		num := strconv.Itoa(int(math.Pow(2, 31)))
		for i := range 9 {
			if num[i] > bytes[i] {
				break
			} else if num[i] < bytes[i] {
				return 0
			}
		}

		if x > 0 && bytes[9] > '7' {
			return 0
		}

		if x < 0 && bytes[9] > '8' {
			return 0
		}
	}

	ans, _ := strconv.Atoi(string(bytes))
	if x < 0 {
		return ans * -1
	}
	return ans
}

func main() {
	x := 123
	fmt.Println(reverse(x))

	x1 := -123
	fmt.Println(reverse(x1))

	x2 := 120
	fmt.Println(reverse(x2))

	x3 := 11111111111
	fmt.Println(reverse(x3))

	x4 := 1000000000
	fmt.Println(reverse(x4))
}
