package main

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func main() {
	n := 4
	fmt.Println(bulbSwitch(n))
}

// func bulbSwitch(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	var res int
// 	for i := 1; i <= n; i++ {
// 		num := int(math.Sqrt(float64(i)))
// 		if i == (num * num) {
// 			res++
// 		}
// 	}

// 	return res
// }
