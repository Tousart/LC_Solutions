package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	n := len(s)
	sMap := make([]int, 26)
	tMap := make([]int, 26)
	for i := range n {
		sMap[s[i]-'a']++
		tMap[t[i]-'a']++
	}

	for i := range 26 {
		if sMap[i] != tMap[i] {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
