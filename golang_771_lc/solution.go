package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[rune]bool)
	count := 0

	for _, val := range jewels {
		jewelsMap[val] = true
	}

	for _, val := range stones {
		if jewelsMap[val] {
			count += 1
		}
	}

	return count
}

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels, stones))
}
