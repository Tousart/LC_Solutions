package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	} else if len(needle) == len(haystack) {
		if needle == haystack {
			return 0
		}
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStr(haystack, needle))

	haystack1 := "mississippi"
	needle1 := "issipi"
	fmt.Println(strStr(haystack1, needle1))
}
