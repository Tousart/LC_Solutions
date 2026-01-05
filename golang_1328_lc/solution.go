package main

import "fmt"

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}

	mid := len(palindrome)
	if mid%2 != 0 {
		mid /= 2
	}

	for i := range palindrome {
		if palindrome[i] > 'a' && i != mid {
			return palindrome[:i] + "a" + palindrome[i+1:]
		}
	}

	return palindrome[:len(palindrome)-1] + "b"
}

func main() {
	s := "abccba"
	fmt.Println(breakPalindrome(s))
}
