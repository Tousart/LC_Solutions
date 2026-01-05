package main

import (
	"fmt"
	"sort"
)

func minDeletions(s string) int {
	var res int

	freqs := make([]int, 26)
	for _, letter := range s {
		freqs[letter-'a']++
	}

	sort.Ints(freqs)

	for i := 24; i >= 0; i-- {
		if freqs[i] == 0 {
			break
		}

		if freqs[i+1] <= freqs[i] {
			curr := freqs[i]
			freqs[i] = max(0, freqs[i+1]-1)
			res += curr - freqs[i]
		}
	}

	return res
}

func main() {
	s := "aab"
	fmt.Println(minDeletions(s))

	s1 := "ceabaacb"
	fmt.Println(minDeletions(s1))

	s2 := "abcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwz"
	fmt.Println(minDeletions(s2))
}

// func minDeletions(s string) int {
// 	var res int

// 	freqs := make([]int, 26)
// 	for _, letter := range s {
// 		freqs[letter-'a']++
// 	}

// 	sort.Ints(freqs)

// 	heights := make([]int, 0)
// 	hasFreq := make(map[int]struct{})
// 	for _, freq := range freqs {
// 		if freq == 0 {
// 			continue
// 		}

// 		var newFreq int

// 		if _, ok := hasFreq[freq]; ok {
// 			if len(heights) > 0 {
// 				diff := freq - heights[len(heights)-1]
// 				heights = heights[:len(heights)-1]
// 				res += diff
// 				hasFreq[freq-diff] = struct{}{}
// 				newFreq = freq - diff - 1
// 			} else {
// 				res += freq
// 			}
// 		} else {
// 			hasFreq[freq] = struct{}{}
// 			newFreq = freq - 1
// 		}

// 		if newFreq > 0 {
// 			if _, ok := hasFreq[newFreq]; !ok {
// 				heights = append(heights, newFreq)
// 			}
// 		}
// 	}

// 	return res
// }
