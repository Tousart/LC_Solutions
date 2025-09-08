package main

import "fmt"

func getNoZeroIntegers(n int) []int {
	n--
	x := 1
	for x < n {
		if !hasZero(x) && !hasZero(n) {
			break
		}
		n--
		x++
	}
	return []int{x, n}
}

func hasZero(num int) bool {
	for num >= 10 {
		if num%10 == 0 {
			return true
		}
		num /= 10
	}
	return false
}

func main() {
	n := 11
	fmt.Println(getNoZeroIntegers(n))
}
