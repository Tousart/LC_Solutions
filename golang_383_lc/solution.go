package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	freqs := make([]int, 26)
	for i := range magazine {
		if i < len(ransomNote) {
			freqs[ransomNote[i]-'a']++
		}
		freqs[magazine[i]-'a']--
	}

	for i := range 26 {
		if freqs[i] > 0 {
			return false
		}
	}

	return true
}

func main() {
	ransomNote := "aa"
	magazine := "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}
