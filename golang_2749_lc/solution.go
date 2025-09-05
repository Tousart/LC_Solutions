package main

import (
	"fmt"
	"math/bits"
)

func makeTheIntegerZero(num1 int, num2 int) int {
	for k := range 61 {
		sumPow := num1 - k*num2

		if sumPow < k {
			continue
		}

		ones := bits.OnesCount64(uint64(sumPow))
		if k >= ones {
			return k
		}
	}
	return -1
}

func main() {
	num1 := 3
	num2 := -2
	fmt.Println(makeTheIntegerZero(num1, num2))
}
