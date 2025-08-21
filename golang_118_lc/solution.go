package main

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{
		{1},
	}

	for i := 1; i < numRows; i++ {
		elem := make([]int, i+1)
		elem[0] = 1
		elem[i] = 1
		for j := 1; j < len(res[i-1]); j++ {
			elem[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, elem)
	}

	return res
}

func main() {
	fmt.Println(generate(5))
}
