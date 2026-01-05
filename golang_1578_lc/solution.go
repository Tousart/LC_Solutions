package main

import "fmt"

func minCost(colors string, neededTime []int) int {
	var res int

	i := 0
	for i < len(colors)-1 {
		if colors[i] == colors[i+1] {
			currMaxTime := neededTime[i]
			sum := neededTime[i]
			j := i + 1
			for j < len(colors) && colors[i] == colors[j] {
				currMaxTime = max(currMaxTime, neededTime[j])
				sum += neededTime[j]
				j++
			}

			res += sum - currMaxTime
			i = j - 1
		}
		i++
	}

	return res
}

func main() {
	colors := "abaac"
	neededTime := []int{1, 2, 3, 4, 5}
	fmt.Println(minCost(colors, neededTime))
}
