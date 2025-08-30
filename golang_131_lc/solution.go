package main

import "fmt"

func partition(s string) [][]string {
	res := [][]string{}
	palik := []string{}
	makePalindrom(s, &res, &palik)
	return res
}

func makePalindrom(s string, res *[][]string, palik *[]string) {
	if s == "" {
		temp := make([]string, len(*palik))
		copy(temp, *palik)
		*res = append(*res, temp)
		return
	}

	for i := 1; i <= len(s); i++ {
		val := s[:i]
		if isPalindrom(val) {
			*palik = append(*palik, val)
			makePalindrom(s[i:], res, palik)
			*palik = (*palik)[:len(*palik)-1]
		}
	}
}

func isPalindrom(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
