package main

import "fmt"

func numSplits(s string) int {
	var (
		res  int
		uniq int
	)

	letters := make([]int, 26)
	for _, letter := range s {
		if letters[letter-'a'] == 0 {
			uniq++
		}
		letters[letter-'a']++
	}

	leftUniq := make(map[rune]struct{})
	for _, letter := range s {
		if _, ok := leftUniq[letter-'a']; !ok {
			leftUniq[letter-'a'] = struct{}{}
		}

		letters[letter-'a']--
		if letters[letter-'a'] == 0 {
			uniq--
		}

		if len(leftUniq) == uniq {
			res++
		} else if len(leftUniq) > uniq {
			return res
		}
	}

	return res
}

func main() {
	s := "aacaba"
	fmt.Println(numSplits(s))
}
