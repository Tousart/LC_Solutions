package main

import (
	"fmt"
)

func simpleComb(n, k int) int {
	if k == n {
		return 1
	}

	if k > n-k {
		k = n - k
	}

	res := 1
	for i := 0; i < k; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}

func countOfOneComb(numCount map[int]int, k int) int {
	res := 1
	for _, count := range numCount {
		res *= simpleComb(count+k-1, k-1)
	}

	return res
}

func makeCombinations(start, sum, n, k int, numCount map[int]int, res *int) {
	if sum == n {
		(*res) += countOfOneComb(numCount, k)
		return
	}

	for i := start; i <= n; i++ {
		if sum+i <= n {
			numCount[i]++
			makeCombinations(i, sum+i, n, k, numCount, res)
			numCount[i]--

			if numCount[i] == 0 {
				delete(numCount, i)
			}
		} else {
			break
		}
	}
}

func main() {
	var (
		n   int
		k   int
		res int
	)
	numCount := make(map[int]int)
	fmt.Scan(&n, &k)

	makeCombinations(1, 0, n, k, numCount, &res)
	fmt.Println(res)
}
