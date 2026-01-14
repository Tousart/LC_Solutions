package main

import "fmt"

func maxProfit(prices []int) int {
	buy1, buy2 := 100001, 100001
	sell1, sell2 := 0, 0

	for _, price := range prices {
		buy1 = min(buy1, price)
		sell1 = max(sell1, price-buy1)
		buy2 = min(buy2, price-sell1)
		sell2 = max(sell2, price-buy2)
	}

	return sell2
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}
