package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var ans []int

	if len(s) < len(p) {
		return ans
	}

	sliceP := make([]int, 26)
	sliceS := make([]int, 26)

	for i := range len(p) {
		sliceP[int(p[i]-'a')]++
		sliceS[int(s[i]-'a')]++
	}

	start := 0
	end := len(p)
	strinP := fmt.Sprint(sliceP)

	if fmt.Sprint(sliceS) == strinP {
		ans = append(ans, start)
	}

	for end < len(s) {
		sliceS[int(s[start]-'a')]--
		sliceS[int(s[end]-'a')]++

		if fmt.Sprint(sliceS) == strinP {
			ans = append(ans, start+1)
		}

		start++
		end++
	}

	return ans
}

func main() {
	s := "baa"
	p := "aa"
	fmt.Println(findAnagrams(s, p))
}
