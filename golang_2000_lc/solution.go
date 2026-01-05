package main

import "fmt"

func reversePrefix(word string, ch byte) string {
	idx := -1
	for i := range word {
		if word[i] == ch {
			idx = i
			break
		}
	}

	if idx == -1 {
		return word
	}

	bytes := []byte(word)
	start := 0
	for start < idx {
		bytes[start], bytes[idx] = bytes[idx], bytes[start]
		start++
		idx--
	}

	return string(bytes)
}

func main() {
	word := "abcdefd"
	ch := 'd'
	fmt.Println(reversePrefix(word, byte(ch)))
}
