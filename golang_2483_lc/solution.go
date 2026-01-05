package main

import "fmt"

func bestClosingTime(customers string) int {
	var (
		allY   int
		countN int
		countY int
		hour   int
	)

	for _, customer := range customers {
		if customer == 'Y' {
			allY++
		}
	}

	if allY == len(customers) {
		return len(customers)
	}

	minPenalty := allY
	for i, customer := range customers {
		if customer == 'N' {
			if minPenalty > (countN + allY - countY) {
				minPenalty = countN + allY - countY
				hour = i
			}
			countN++
		} else {
			countY++
		}
	}

	if minPenalty > countN {
		return len(customers)
	}

	return hour
}

func main() {
	customers := "YYNY"
	fmt.Println(bestClosingTime(customers))

	customers1 := "YNYY"
	fmt.Println(bestClosingTime(customers1))
}
