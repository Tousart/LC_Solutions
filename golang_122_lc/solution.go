package main

import "fmt"

func maxProfit(prices []int) int {
	var res int

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

func main() {
	prices := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))

	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices1))
}
