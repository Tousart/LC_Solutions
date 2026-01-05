package main

import "fmt"

func largestMerge(word1 string, word2 string) string {
	idx1, idx2, idxRes := 0, 0, 0
	res := make([]byte, len(word1)+len(word2))

	for idx1 < len(word1) && idx2 < len(word2) {
		if word1[idx1] > word2[idx2] {
			res[idxRes] = word1[idx1]
			idx1++
		} else if word1[idx1] < word2[idx2] {
			res[idxRes] = word2[idx2]
			idx2++
		} else {
			if word1[idx1:] > word2[idx2:] {
				res[idxRes] = word1[idx1]
				idx1++
			} else {
				res[idxRes] = word2[idx2]
				idx2++
			}
		}
		idxRes++
	}

	for idx1 < len(word1) {
		res[idxRes] = word1[idx1]
		idx1++
		idxRes++
	}

	for idx2 < len(word2) {
		res[idxRes] = word2[idx2]
		idx2++
		idxRes++
	}

	return string(res)
}

func main() {
	word1 := "cabaa"
	word2 := "bcaaa"
	fmt.Println(largestMerge(word1, word2))
}
