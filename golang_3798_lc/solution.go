package main

import "fmt"

func largestEven(s string) string {
	idx := len(s) - 1
	for idx >= 0 {
		if s[idx] == '2' {
			break
		}
		idx--
	}
	return s[:idx+1]
}

func main() {
	s := "121"
	fmt.Println(largestEven(s))
}
