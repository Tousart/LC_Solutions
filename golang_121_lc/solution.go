package main

import "fmt"

func maxProfit(prices []int) int {
	minLeft := 10001
	res := 0
	for _, price := range prices {
		if price > minLeft {
			res = max(res, price-minLeft)
		} else {
			minLeft = price
		}
	}
	return res
}

func main() {
	prices := []int{2, 1, 3, 4}
	fmt.Println(maxProfit(prices))
}
