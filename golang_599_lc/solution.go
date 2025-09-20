package main

import (
	"fmt"
)

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		return findRestaurant(list2, list1)
	}

	words := make(map[string]int)
	for i, val := range list1 {
		words[val] = i
	}

	res := []string{}
	minSum := len(list1) + len(list2)
	for i, val := range list2 {
		if ind, ok := words[val]; ok {
			sum := ind + i
			if sum < minSum {
				res = []string{val}
				minSum = sum
			} else if sum == minSum {
				res = append(res, val)
			} else {
				break
			}
		}
	}

	return res
}

func main() {
	list1 := []string{"huy", "priv", "sasi"}
	list2 := []string{"priv", "huy", "sasi"}
	fmt.Println(findRestaurant(list1, list2))
}
