package main

import (
	"fmt"
	"sort"
)

func minimumBoxes(apple []int, capacity []int) int {
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})

	i, j := 0, 0
	packageInUse := false
	for i < len(apple) {
		if apple[i] < capacity[j] {
			capacity[j] -= apple[i]
			i++
			packageInUse = true
		} else if apple[i] > capacity[j] {
			apple[i] -= capacity[j]
			j++
			packageInUse = true
		} else {
			i++
			j++
			packageInUse = false
		}
	}

	if !packageInUse {
		return j
	}

	return j + 1
}

func main() {
	apple := []int{1, 3, 2}
	capacity := []int{4, 3, 1, 5, 2}
	fmt.Println(minimumBoxes(apple, capacity))

	apple1 := []int{3}
	capacity1 := []int{2, 3}
	fmt.Println(minimumBoxes(apple1, capacity1))
}
