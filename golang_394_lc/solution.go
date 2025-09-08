package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	res := []byte{}
	i := 0
	for i < len(s) {
		if unicode.IsDigit(rune(s[i])) {
			j := i + 1
			for s[j] != '[' {
				j++
			}
			num, _ := strconv.Atoi(s[i:j])

			count := 1
			k := j + 1
			for k < len(s) && count != 0 {
				switch s[k] {
				case ']':
					count--
				case '[':
					count++
				}
				k++
			}

			res = append(res, strings.Repeat(decodeString(s[j+1:k-1]), num)...)
			i = k
		} else if s[i] == ']' {
			i++
		} else {
			res = append(res, s[i])
			i++
		}
	}
	return string(res)
}

func main() {
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))

	s1 := "3[a]2[b2[dd]]"
	fmt.Println(decodeString(s1))
}
