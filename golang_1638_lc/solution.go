package main

import "fmt"

func countSubstrings(s string, t string) int {
	var res int

	for i := range s {
		for j := range t {
			miss := 0
			idx := 0
			for i+idx < len(s) && j+idx < len(t) {
				if s[i+idx] != t[j+idx] {
					miss++
					if miss == 2 {
						break
					}
				}
				res += miss
				idx++
			}
		}
	}

	return res
}

func main() {
	s := "aba"
	t := "baba"
	fmt.Println(countSubstrings(s, t))
}
