package main

import "fmt"

func countOdds(low int, high int) int {
	if (low == high) && (low%2 == 0) {
		return 0
	}

	if low%2 == 0 {
		low++
	}

	if high%2 == 0 {
		high--
	}

	return ((high - low + 1) / 2) + 1
}

func main() {
	low := 15
	high := 22
	fmt.Println(countOdds(low, high))
}
