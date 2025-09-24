package main

import "fmt"

func countSubstrings(s string) int {
	var res int
	for i := range s {
		res += countOfPalindroms(s, i, i) + countOfPalindroms(s, i, i+1)
	}
	return res
}

func countOfPalindroms(s string, l, r int) int {
	var count int
	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}
	return count
}

func main() {
	s := "aaa"
	fmt.Println(countSubstrings(s))
}
