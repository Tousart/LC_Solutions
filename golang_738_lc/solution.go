package main

import (
	"fmt"
)

func monotoneIncreasingDigits(n int) int {
	shift := 1
	prev := 9
	for shift <= n {
		num := n / shift % 10
		if num <= prev {
			prev = num
		} else {
			n = n/shift*shift - 1
			prev = n / shift % 10
		}
		shift *= 10
	}
	return n
}

func main() {
	n := 332
	fmt.Println(monotoneIncreasingDigits(n))

	n1 := 110
	fmt.Println(monotoneIncreasingDigits(n1))

	n2 := 10
	fmt.Println(monotoneIncreasingDigits(n2))
}
