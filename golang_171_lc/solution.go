package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {
	var res int
	for i := range len(columnTitle) {
		res = int((columnTitle[i] - 'A')) + 1 + res*26
	}
	return res
}

func main() {
	col := "ZY"
	fmt.Println(titleToNumber(col))
}
