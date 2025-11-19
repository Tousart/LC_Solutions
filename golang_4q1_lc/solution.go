package main

import "fmt"

func finalPrices(prices []int) []int {
	needDisc := make([][2]int, 0)

	for i, price := range prices {
		for len(needDisc) > 0 && needDisc[len(needDisc)-1][1] >= price {
			idxVal := needDisc[len(needDisc)-1]
			prices[idxVal[0]] = idxVal[1] - price
			needDisc = needDisc[:len(needDisc)-1]
		}

		needDisc = append(needDisc, [2]int{i, price})
	}

	return prices
}

func main() {
	prices := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices(prices))
}
