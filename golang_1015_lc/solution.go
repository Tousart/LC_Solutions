package main

import "fmt"

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	remain := 0
	for length := 1; length <= k; length++ {
		remain = (remain*10 + 1) % k
		if remain == 0 {
			return length
		}
	}

	return -1
}

func main() {
	k := 1
	fmt.Println(smallestRepunitDivByK(k))

	k1 := 2
	fmt.Println(smallestRepunitDivByK(k1))

	k2 := 37
	fmt.Println(smallestRepunitDivByK(k2))
}
