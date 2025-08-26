package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	runes := []rune(strings.TrimLeft(s, " "))
	if len(runes) == 0 {
		return 0
	}

	if unicode.IsDigit(runes[0]) {
		res, _ := strconv.Atoi(makeNum(runes, '+'))
		return res
	} else if runes[0] == '+' {
		res, _ := strconv.Atoi(makeNum(runes[1:], '+'))
		return res
	} else if runes[0] == '-' {
		res, _ := strconv.Atoi("-" + makeNum(runes[1:], '-'))
		return res
	}

	return 0
}

func makeNum(runes []rune, sign rune) string {
	start, end := -1, -1
	for i, num := range runes {
		if !unicode.IsDigit(num) {
			break
		} else if start == -1 {
			if num == '0' {
				continue
			}
			start = i
		}
		end = i
	}

	if start == -1 {
		return "0"
	}

	n := end - start + 1

	var border []rune
	if sign == '+' {
		border = []rune("2147483647")
	} else {
		border = []rune("2147483648")
	}

	if n > len(border) {
		return string(border)
	} else if n == len(border) {
		for i := range border {
			if runes[i+start] > border[i] {
				return string(border)
			} else if runes[i+start] < border[i] {
				return string(runes[start : end+1])
			}
		}
	}

	return string(runes[start : end+1])
}

func main() {
	s := "21474836460"
	fmt.Println(myAtoi(s))

	s1 := "a"
	fmt.Println(myAtoi(s1))

	s2 := "   -42"
	fmt.Println(myAtoi(s2))

	s3 := "0-2"
	fmt.Println(myAtoi(s3))

	s4 := "3.14159"
	fmt.Println(myAtoi(s4))
}
