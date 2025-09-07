package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	seqMap := make(map[string]bool)
	for i := 0; i <= len(s)-10; i++ {
		seq := s[i : i+10]
		if _, ok := seqMap[seq]; !ok {
			seqMap[seq] = false
		} else {
			if !seqMap[seq] {
				res = append(res, seq)
				seqMap[seq] = true
			}
		}
	}
	return res
}

func main() {
	seq := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(seq))
}
