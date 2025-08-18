package main

import (
	"fmt"
	"math"
	"strings"
)

func firstUniqChar(s string) int {
	res := math.MaxInt32
	for i := 'a'; i <= 'z'; i++ {
		if strings.Contains(s, string(i)) {
			let := string(i)
			if strings.Index(s, let) == strings.LastIndex(s, let) {
				res = minimum(res, strings.Index(s, let))
			}
		}
	}
	if res != math.MaxInt32 {
		return res
	}
	return -1
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "lllllleetcode"
	fmt.Println(firstUniqChar(s))
}

// for i, val := range s {
//     if strings.Index(s, string(val)) == strings.LastIndex(s, string(val)) {
//         return i
//     }
// }

// func firstUniqChar(s string) int {
// 	letters := make(map[rune]int)
// 	for _, val := range s {
// 		letters[val]++
// 	}

// 	for i, val := range s {
// 		if letters[val] == 1 {
// 			return i
// 		}
// 	}

// 	return -1
// }
