package main

import (
	"fmt"
)

func reverseWords(s string) string {
	bytes := []byte(s)
	start, j, k := -1, 0, 0
	for i := 0; i <= len(bytes); i++ {
		if i == len(bytes) || bytes[i] == ' ' {
			start++
			j = start
			k = i - 1
			for j < k {
				bytes[j], bytes[k] = bytes[k], bytes[j]
				j++
				k--
			}
			start = i
		}
	}
	return string(bytes)
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
