package main

import "fmt"

func findTheDifference(s string, t string) byte {
	sumS := byte(0)
	sumT := t[len(s)]
	for i := range len(s) {
		sumS += s[i]
		sumT += t[i]
	}
	return sumT - sumS
}

func main() {
	s := "abc"
	t := "bcae"
	fmt.Println(findTheDifference(s, t))
}
