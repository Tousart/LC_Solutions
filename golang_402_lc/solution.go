package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	if k == 0 {
		return num
	}

	res := make([]rune, 0)

	for _, n := range num {
		for len(res) > 0 && res[len(res)-1] > n && k > 0 {
			res = res[:len(res)-1]
			k--
		}

		if len(res) > 0 || n != '0' {
			res = append(res, n)
		}
	}

	for k > 0 && len(res) > 0 {
		res = res[:len(res)-1]
		k--
	}

	if len(res) == 0 {
		return "0"
	}

	return string(res)
}

func main() {
	num := "10199"
	k := 2
	fmt.Println(removeKdigits(num, k))
}
