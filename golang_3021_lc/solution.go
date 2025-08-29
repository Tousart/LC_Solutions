package main

import "fmt"

func flowerGame(n int, m int) int64 {
	return int64(n * m / 2)
}

func main() {
	n := 3
	m := 2
	fmt.Println(flowerGame(n, m))
}
