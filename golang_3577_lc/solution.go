package main

import "fmt"

func countPermutations(complexity []int) int {
	res := 1
	mod := 1000000007

	for i := 1; i < len(complexity); i++ {
		if complexity[0] >= complexity[i] {
			return 0
		}
		res = (i * res) % mod
	}

	return res
}

func main() {
	complexity := []int{1, 2, 3}
	fmt.Println(countPermutations(complexity))
}
