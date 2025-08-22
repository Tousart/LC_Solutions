package main

import (
	"fmt"
	"math"
)

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	res := len(fruits)
	for i := range fruits {
		for j := range baskets {
			if fruits[i] <= baskets[j] {
				baskets[j] = math.MinInt64
				res--
				break
			}
		}
	}
	return res
}

func main() {
	fruits := []int{4, 2, 5}
	baskets := []int{3, 5, 4}
	fmt.Println(numOfUnplacedFruits(fruits, baskets))
}
