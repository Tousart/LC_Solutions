package main

import (
	"fmt"
	"math"
)

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	x := math.Log(float64(n)) / math.Log(3.0)
	return math.Abs(math.Round(x)-x) < 1e-10
}

func main() {
	n := 27
	fmt.Println(isPowerOfThree(n))
	n1 := 45
	fmt.Println(isPowerOfThree(n1))
}
