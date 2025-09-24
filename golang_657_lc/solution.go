package main

import "fmt"

func judgeCircle(moves string) bool {
	upDown := 0
	rightLeft := 0
	for _, move := range moves {
		switch move {
		case 'U':
			upDown++
		case 'D':
			upDown--
		case 'R':
			rightLeft++
		case 'L':
			rightLeft--
		}
	}
	return upDown == 0 && rightLeft == 0
}

func main() {
	moves := "UDLR"
	fmt.Println(judgeCircle(moves))
}
