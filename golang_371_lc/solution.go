package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a ^= b
		b = carry << 1
	}
	return a
}

//  func getSum1(a int, b int) int {
//  	return int(math.Log(math.Exp(float64(a)) * math.Exp(float64(b))))
// }

func main() {
	a := -999
	b := 0
	fmt.Println(getSum(a, b))
	// fmt.Println(getSum1(a, b))
}
