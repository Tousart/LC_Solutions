package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func searchWinner(nums []string) string {
	slices.Sort(nums)
	sumP := 0
	sumA := 0
	for i := 1; i <= len(nums); i++ {
		elem, _ := strconv.Atoi(nums[i-1])
		if elem > i {
			return "Second"
		}
		sumP += i
		sumA += elem
	}

	if (sumP-sumA)%2 == 0 {
		return "Second"
	}
	return "First"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for range n {
		scanner.Scan()
		scanner.Scan()
		fmt.Println(searchWinner(strings.Fields(scanner.Text())))
	}
}
