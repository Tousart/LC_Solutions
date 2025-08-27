package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(prefix) {
			prefix = strs[i]
		}
	}

	end := len(prefix)
	for _, word := range strs {
		for j := 0; j < end; j++ {
			if prefix[j] != word[j] {
				end = j
				if end < 0 {
					return ""
				}
			}
		}
	}

	return prefix[:end]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
