package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	var res int

	for i, ticket := range tickets {
		if i == k {
			res += ticket
		} else if i < k {
			res += min(tickets[k], ticket)
		} else {
			res += min(tickets[k]-1, ticket)
		}
	}

	return res
}

func main() {
	tickets := []int{2, 3, 2}
	k := 2
	fmt.Println(timeRequiredToBuy(tickets, k))
}
