package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortVowels(s string) string {
	vowels := "aeiouAEIOU"
	vv := []rune{}
	t := []rune(s)

	for _, val := range s {
		if strings.ContainsRune(vowels, val) {
			vv = append(vv, val)
		}
	}

	sort.Slice(vv, func(i, j int) bool {
		return vv[i] < vv[j]
	})

	for i, val := range s {
		if strings.ContainsRune(vowels, val) {
			t[i] = vv[0]
			vv = vv[1:]
		}
	}
	return string(t)
}

func main() {
	s := "lEetcOde"
	fmt.Println(sortVowels(s))

	s1 := "lYmpH"
	fmt.Println(sortVowels(s1))

	s2 := "axRukCyOHm"
	fmt.Println(sortVowels(s2))
}
