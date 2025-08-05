package main

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var (
		result [][]int
		first  int
		second int
		minR   int
	)

	for first < len(firstList) && second < len(secondList) {
		fL := firstList[first][0]
		fR := firstList[first][1]
		sL := secondList[second][0]
		sR := secondList[second][1]

		if fR <= sR {
			minR = fR
			first++
		} else {
			minR = sR
			second++
		}

		if sL > fR || fL > sR {
			continue
		}

		result = append(result, []int{maximum(fL, sL), minR})
	}

	return result
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	first := [][]int{
		{0, 2},
		{5, 10},
	}

	second := [][]int{
		{1, 5},
		{8, 12},
	}

	result := intervalIntersection(first, second)
	for i := 0; i < len(result); i++ {
		fmt.Println()
		for _, val := range result[i] {
			fmt.Print(val, " ")
		}
	}
}
