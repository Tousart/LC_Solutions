package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := make([][]int, 0)
	combination := make([]int, 0)

	var makeCombinations func(idx, t int)
	makeCombinations = func(idx, t int) {
		if t == 0 {
			comb := make([]int, len(combination))
			copy(comb, combination)
			res = append(res, comb)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if candidates[i] > t {
				break
			}

			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}

			combination = append(combination, candidates[i])
			makeCombinations(i+1, t-candidates[i])
			combination = combination[:len(combination)-1]
		}
	}

	makeCombinations(0, target)

	return res
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
