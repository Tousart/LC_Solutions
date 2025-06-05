package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var (
		ans   [][]int = [][]int{intervals[0]}
		elemI []int
		elemA []int
	)

	for i := 1; i < len(intervals); i++ {
		elemI = intervals[i]
		elemA = ans[len(ans)-1]
		if elemI[0] >= elemA[0] && elemI[0] <= elemA[1] {
			ans[len(ans)-1] = []int{min(elemI[0], elemA[0]), max(elemI[1], elemA[1])}
		} else {
			ans = append(ans, elemI)
		}
	}
	return ans
}

func main() {
	intervals := [][]int{{1, 4}, {0, 7}}
	fmt.Println(merge(intervals))
}
