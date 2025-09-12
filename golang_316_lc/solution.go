package main

import (
	"fmt"
)

func removeDuplicateLetters(s string) string {
	resStack := []rune{}
	lastInds := make(map[rune]int)
	inStack := make(map[rune]struct{})
	for i, val := range s {
		lastInds[val] = i
	}

	for i, val := range s {
		if _, ok := inStack[val]; !ok {
			for len(resStack) > 0 && resStack[len(resStack)-1] > val && lastInds[resStack[len(resStack)-1]] > i {
				delete(inStack, resStack[len(resStack)-1])
				resStack = resStack[:len(resStack)-1]
			}
			inStack[val] = struct{}{}
			resStack = append(resStack, val)
		}
	}

	return string(resStack)
}

func main() {
	s := "bcabc"
	fmt.Println(removeDuplicateLetters(s))
}
