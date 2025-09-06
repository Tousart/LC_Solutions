package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	sliceS := []string{}
	i, j := 0, 0
	for i < len(s) {
		if s[i] != ' ' {
			j = i + 1
			for j < len(s) {
				if s[j] == ' ' {
					break
				}
				j++
			}
			sliceS = append(sliceS, s[i:j])
			i = j
		}
		i++
	}

	i, j = 0, len(sliceS)-1
	for i < j {
		sliceS[i], sliceS[j] = sliceS[j], sliceS[i]
		i++
		j--
	}
	return strings.Join(sliceS, " ")
}

func main() {
	s := "sf   fff    "
	fmt.Println(reverseWords(s))
}
