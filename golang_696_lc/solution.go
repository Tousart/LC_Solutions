package main

import "fmt"

func countBinarySubstrings(s string) int {
	res, i, j := 0, 0, 0
	for j < len(s) {
		if s[j] != s[i] {
			pivot := j
			for j < len(s) {
				if s[j] == s[i] {
					break
				}
				j++
			}
			res += min(pivot-i, j-pivot)
			i = pivot
		} else {
			j++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "00110011"
	fmt.Println(countBinarySubstrings(s))
}
