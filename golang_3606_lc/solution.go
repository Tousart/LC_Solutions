package main

import (
	"fmt"
	"sort"
	"strconv"
)

var coupones = map[string]int{
	"electronics": 1,
	"grocery":     2,
	"pharmacy":    3,
	"restaurant":  4,
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	toSort := make([][]string, 0)

	for i, c := range code {
		if !isActive[i] {
			continue
		}

		if _, ok := coupones[businessLine[i]]; !ok {
			continue
		}

		if len(c) == 0 || !compare(c) {
			continue
		}

		toSort = append(toSort, []string{c, strconv.Itoa(i)})
	}

	sort.Slice(toSort, func(i, j int) bool {
		idxI, _ := strconv.Atoi(toSort[i][1])
		idxJ, _ := strconv.Atoi(toSort[j][1])
		if coupones[businessLine[idxI]] == coupones[businessLine[idxJ]] {
			return toSort[i][0] < toSort[j][0]
		}
		return coupones[businessLine[idxI]] < coupones[businessLine[idxJ]]
	})

	res := make([]string, len(toSort))

	for i, pair := range toSort {
		res[i] = pair[0]
	}

	return res
}

func compare(code string) bool {
	for _, symbol := range code {
		if !(symbol >= 'A' && symbol <= 'Z') &&
			!(symbol >= 'a' && symbol <= 'z') &&
			!(symbol >= '0' && symbol <= '9') &&
			!(symbol == '_') {
			return false
		}
	}
	return true
}

func main() {
	code := []string{"m", "A", "B", "P", "J", "P", "u", "W", "4", "J", "C", "9"}
	businessLine := []string{"electronics", "invalid", "invalid", "pharmacy", "invalid", "electronics", "restaurant", "grocery", "restaurant", "pharmacy", "pharmacy", "pharmacy"}
	isActive := []bool{true, true, false, true, false, true, true, false, false, false, true, false}
	fmt.Println(validateCoupons(code, businessLine, isActive))
}
