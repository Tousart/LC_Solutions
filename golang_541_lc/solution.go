package main

import "fmt"

func reverseStr(s string, k int) string {
	res := make([]byte, len(s))
	for i := 0; i < len(s); i += 2 * k {
		l := i + k - 1
		r := i + k
		idx := i
		for l >= i {
			if l < len(s) {
				res[idx] = s[l]
				idx++
			}
			if r < len(s) {
				res[r] = s[r]
			}
			l--
			r++
		}
	}
	return string(res)
}

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseStr(s, k))
}
