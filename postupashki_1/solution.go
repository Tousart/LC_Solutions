package main

import (
	"fmt"
	"strings"
)

func changeOrder(s string) string {
	words := []string{}
	i := 0
	for i < len(s) {
		if s[i] != ' ' {
			j := i + 1
			for j < len(s) {
				if s[j] == ' ' {
					words = append(words, s[i:j])
					i = j
					break
				}
				j++
			}

			if j == len(s) {
				words = append(words, s[i:j])
				break
			}
		}
		i++
	}

	if len(words) == 0 || len(words[0]) == len(s) {
		return s
	}

	ind := len(words) - 1

	res := []string{}
	start, end := 0, 0
	for end < len(s) {
		if s[end] != ' ' {
			if start != end {
				res = append(res, s[start:end])
			}

			res = append(res, words[ind])
			ind--

			j := end + 1
			for j < len(s) {
				if s[j] == ' ' {
					start = j
					end = j
					break
				}
				j++
			}

			if j == len(s) {
				start = j
				end = j
				break
			}
		}
		end++
	}

	if start != end {
		res = append(res, s[start:end])
	}

	return strings.Join(res, "")
}

func main() {
	s := "  hello my   dear world "
	fmt.Println(changeOrder(s))

	s1 := "hello my   dear world "
	fmt.Println(changeOrder(s1))

	s2 := "   hello my   dear world"
	fmt.Println(changeOrder(s2))

	s3 := " "
	fmt.Println(changeOrder(s3))
}
