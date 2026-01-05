package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	var (
		res   int
		count int
	)

	for i := range k {
		if count >= happiness[i] {
			break
		}
		res += happiness[i] - count
		count++
	}

	return int64(res)
}

func main() {
	happiness := []int{1, 2, 3}
	k := 2
	fmt.Println(maximumHappinessSum(happiness, k))
}
