package main

import (
	"fmt"
	"math/rand/v2"
)

func rand10() int {
	val := 41
	for val > 40 {
		val = (rand7()-1)*7 + rand7()
	}
	return val%10 + 1
}
func rand7() int {
	return rand.IntN(7-1+1) + 1
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(rand10())
	}
}
