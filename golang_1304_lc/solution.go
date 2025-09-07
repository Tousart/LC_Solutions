package main

import "fmt"

func sumZero(n int) []int {
	res := make([]int, n)
	mid := n / 2
	k := 0

	if n%2 == 0 {
		res[mid] = 1
		res[mid-1] = -1
		k = 2
	} else {
		res[mid] = 0
		k = 1
	}

	for i := mid + 1; i < n; i++ {
		res[i] = res[i-1] + 1
	}

	for i := mid - k; i >= 0; i-- {
		res[i] = res[i+1] - 1
	}

	return res
}

func main() {
	n := 4
	fmt.Println(sumZero(n))

	n1 := 5
	fmt.Println(sumZero(n1))
}
