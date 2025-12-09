package main

import (
	"fmt"
	"math"
)

func countTriples(n int) int {
	var res int

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			target := int(math.Sqrt(float64(i*i + j*j)))
			if (target*target) == (i*i+j*j) && target <= n {
				if i == j {
					res++
				} else {
					res += 2
				}
			}
		}
	}

	return res
}

func main() {
	n := 5
	fmt.Println(countTriples(n))
}
