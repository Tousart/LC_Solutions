package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	step := len(s1)
	s1Freq := make([]int, 26)
	s2Freq := make([]int, 26)

	for i := 0; i < step; i++ {
		s1Freq[s1[i]-'a']++
		s2Freq[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-step; i++ {
		if compare(s1Freq, s2Freq) {
			return true
		}

		s2Freq[s2[i]-'a']--
		s2Freq[s2[i+step]-'a']++
	}
	return compare(s1Freq, s2Freq)
}

func compare(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := "ab"
	s2 := "cab"
	fmt.Println(checkInclusion(s1, s2))
}
