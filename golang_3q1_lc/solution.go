package main

import "fmt"

func buildArray(target []int, n int) []string {
	res := make([]string, 0)
	tIdx := 0
	stack := make([]int, len(target))
	for i := range n {
		stack[tIdx] = i + 1
		res = append(res, "Push")
		if i+1 != target[tIdx] {
			res = append(res, "Pop")
		} else {
			tIdx++
		}
		if tIdx == len(target) {
			break
		}
	}
	return res
}

func main() {
	target := []int{1, 3}
	n := 3
	fmt.Println(buildArray(target, n))
}
