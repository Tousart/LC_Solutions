package main

import "fmt"

func reverseOnlyLetters(s string) string {
	res := make([]byte, len(s))
	l, r := 0, len(s)-1
	for l <= r {
		if (s[l] > 'Z' && s[l] < 'a') || s[l] < 'A' || s[l] > 'z' {
			res[l] = s[l]
			l++
			continue
		}

		if (s[r] > 'Z' && s[r] < 'a') || s[r] < 'A' || s[r] > 'z' {
			res[r] = s[r]
			r--
			continue
		}

		res[l] = s[r]
		res[r] = s[l]
		l++
		r--
	}
	return string(res)
}

func main() {
	s := "a-bC-dEf-ghIj"
	fmt.Println(reverseOnlyLetters(s))

	s1 := "ab-cd"
	fmt.Println(reverseOnlyLetters(s1))
}
