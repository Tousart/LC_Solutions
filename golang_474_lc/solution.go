package main

import (
	"fmt"
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	limits := make([][]int, m+1)
	for i := range limits {
		limits[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeroes := strings.Count(str, "0")
		ones := len(str) - zeroes
		for i := m; i >= zeroes; i-- {
			for j := n; j >= ones; j-- {
				count := limits[i-zeroes][j-ones] + 1
				if count > limits[i][j] {
					limits[i][j] = count
				}
			}
		}
	}

	return limits[m][n]
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3
	fmt.Println(findMaxForm(strs, m, n))
}
