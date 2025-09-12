package main

import (
	"fmt"
)

func doesAliceWin(s string) bool {
	for _, val := range s {
		if 0x104111>>(val-'a')&1 != 0 {
			return true
		}
	}
	return false
}

func main() {
	s := "sdsfsgge"
	fmt.Println(doesAliceWin(s))
}
