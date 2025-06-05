package main

import (
	"fmt"
)

func maxDistToClosest(seats []int) int {
	ans := 0
	last := -1

	for i := range seats {
		if seats[i] == 1 {
			if last < 0 {
				ans = i
			} else {
				ans = max(ans, (i-last)/2)
			}
			last = i
		}
	}
	return max(ans, len(seats)-1-last)
}

func main() {
	seats := []int{0, 1}
	fmt.Println(maxDistToClosest(seats))
}
