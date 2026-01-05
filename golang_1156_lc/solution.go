package main

import "fmt"

func maxRepOpt1(text string) int {
	res := 1

	for let := 'a'; let <= 'z'; let++ {
		i, j := 0, 0
		k := 1
		letCount := 0
		for j < len(text) {
			if text[j] != byte(let) {
				k--
			} else {
				letCount++
			}
			j++

			if k < 0 {
				if text[i] != byte(let) {
					k++
				}
				i++
			}
		}
		res = max(res, min(j-i, letCount))
	}

	return res
}

func main() {
	text := "aaabaaabbbaaac"
	fmt.Println(maxRepOpt1(text))

	text1 := "aaabaaa"
	fmt.Println(maxRepOpt1(text1))

	text2 := "aaabbba"
	fmt.Println(maxRepOpt1(text2))
}
