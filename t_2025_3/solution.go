package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func count(nums []string) int {
	res := 1
	numsCount := make(map[string]int)
	for _, val := range nums {
		numsCount[val]++
	}

	for _, val := range numsCount {
		res *= val + 1
	}

	return (res - 1) % 1000000007
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	fmt.Println(count(strings.Fields(scanner.Text())))
}
