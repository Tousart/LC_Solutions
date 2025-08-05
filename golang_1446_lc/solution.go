package main

import "fmt"

func maxPower(s string) int {
	var (
		ans   int
		i     int
		j     int = 1
		power int = 1
	)

	for i < len(s)-1 && j < len(s) {
		if s[i] == s[j] {
			power++
		} else {
			ans = maximum(ans, power)
			i = j
			power = 1
		}
		j++
	}

	ans = maximum(ans, power)

	return ans
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "aaabbcccc"
	fmt.Println(maxPower(s))
}
