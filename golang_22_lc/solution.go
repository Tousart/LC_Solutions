package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	makeCombinations(n, &result, 1, 0, "(")
	return result
}

func makeCombinations(n int, result *[]string, left int, right int, combination string) {
	if left == n && right == n {
		*result = append(*result, combination)
		return
	} else if left > n || right > n || left < right {
		return
	}

	makeCombinations(n, result, left+1, right, combination+"(")
	makeCombinations(n, result, left, right+1, combination+")")
}

func main() {
	fmt.Println(generateParenthesis(3))
}
