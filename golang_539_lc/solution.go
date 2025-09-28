package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	res := math.MaxInt64
	minutes := make([]int, len(timePoints))
	for i, timePoint := range timePoints {
		h, _ := strconv.Atoi(timePoint[0:2])
		m, _ := strconv.Atoi(timePoint[3:])
		minutes[i] = 60*h + m
	}

	sort.Ints(minutes)

	for i := 0; i < len(minutes)-1; i++ {
		res = min(res, minutes[i+1]-minutes[i])
		if res == 0 {
			return res
		}
	}

	return min(res, 1440-minutes[len(minutes)-1]+minutes[0])
}

func main() {
	timePoints := []string{
		"23:59",
		"00:00",
	}
	fmt.Println(findMinDifference(timePoints))
}
