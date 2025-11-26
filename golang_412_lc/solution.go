package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := make([]string, n+1)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res[i] = "FizzBuzz"
		} else if i%3 == 0 {
			res[i] = "Fizz"
		} else if i%5 == 0 {
			res[i] = "Buzz"
		} else {
			res[i] = strconv.Itoa(i)
		}
	}

	return res[1:]
}

func main() {
	n := 3
	fmt.Println(fizzBuzz(n))
}
