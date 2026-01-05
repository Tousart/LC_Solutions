package main

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool {
	if len(goal) != len(s) {
		return false
	}
	return strings.Contains(s+s, goal)
}

func main() {
	s := "abcde"
	goal := "cdeab"
	fmt.Println(rotateString(s, goal))
}
