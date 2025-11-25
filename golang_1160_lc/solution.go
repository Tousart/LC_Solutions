package main

import "fmt"

func countCharacters(words []string, chars string) int {
	var res int
	freq := makeFreq(chars)

	for _, word := range words {
		currFreq := makeFreq(word)
		i := 0
		for i < 26 {
			if currFreq[i] > freq[i] {
				break
			}
			i++
		}
		if i == 26 {
			res += len(word)
		}
	}

	return res
}

func makeFreq(word string) []int {
	freq := make([]int, 26)

	for _, letter := range word {
		freq[letter-'a']++
	}

	return freq
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Println(countCharacters(words, chars))
}
