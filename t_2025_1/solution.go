package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func swapNums(inputNum []rune) int {
	slices.Sort(inputNum)

	if inputNum[0] == '0' {
		if inputNum[1] != '0' {
			inputNum[0], inputNum[1] = inputNum[1], inputNum[0]
		} else if inputNum[2] != '0' {
			inputNum[0], inputNum[2] = inputNum[2], inputNum[0]
		} else {
			inputNum[0], inputNum[3] = inputNum[3], inputNum[0]
		}
	}

	ans, _ := strconv.Atoi(string(inputNum))
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Print(swapNums([]rune(scanner.Text())))
}
