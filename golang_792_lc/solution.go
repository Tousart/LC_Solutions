package main

import (
	"fmt"
)

func numMatchingSubseq(s string, words []string) int {
	var res int

	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}

	for word, count := range wordsMap {
		if isSubsequence(s, word) {
			res += count
		}
	}

	return res
}

func isSubsequence(s, word string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(word) {
		if word[j] == s[i] {
			j++
		}
		i++
	}
	return j == len(word)
}

func main() {
	s := "abcde"
	words := []string{"a", "bb", "ace", "acd"}
	fmt.Println(numMatchingSubseq(s, words))
}

// beats 66.67% time

// func numMatchingSubseq(s string, words []string) int {
// 	var res int

// 	inds := make([][]int, 26)
// 	for i, letter := range s {
// 		inds[letter-'a'] = append(inds[letter-'a'], i)
// 	}

// 	for _, word := range words {
// 		idx := -1
// 		flag := true

// 		for _, letter := range word {
// 			slice := inds[letter-'a']

// 			if slice == nil || slice[len(slice)-1] < idx {
// 				flag = false
// 				break
// 			}

// 			l, r := 0, len(slice)-1
// 			for l < r {
// 				mid := (l + r) / 2
// 				if slice[mid] >= idx {
// 					r = mid
// 				} else {
// 					l = mid + 1
// 				}
// 			}
// 			idx = slice[l] + 1
// 		}

// 		if flag {
// 			res++
// 		}
// 	}

// 	return res
// }
