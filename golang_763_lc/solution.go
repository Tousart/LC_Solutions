package main

import "fmt"

func partitionLabels(s string) []int {
	res := []int{}

	inds := make(map[rune]int)
	for i, val := range s {
		inds[val] = i
	}

	start, end := 0, 0
	for i, val := range s {
		end = maximum(end, inds[val])
		if i == end {
			res = append(res, end-start+1)
			end = i + 1
			start = end
		}
	}

	return res
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}
