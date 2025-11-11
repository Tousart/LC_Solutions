package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minRoute(a, b, c, v0, v1, v2 float64) float64 {
	t1 := ((a + b) / v0) + ((a + b) / v1)
	t2 := (a / v0) + (c / v1) + (b / v2)
	t3 := (b / v0) + (c / v1) + (a / v2)
	t4 := (a / v0) + (c / v0) + (c / v1) + (a / v1) + (a / v0) + (a / v1)
	t5 := (b / v0) + (c / v0) + (c / v1) + (b / v1) + (b / v0) + (b / v1)
	t6 := (a / v0) + (c / v0) + (c / v1) + (a / v2)
	t7 := (b / v0) + (c / v0) + (c / v1) + (b / v2)
	return min(t1, t2, t3, t4, t5, t6, t7)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	a, _ := strconv.ParseFloat(input[0], 64)
	b, _ := strconv.ParseFloat(input[1], 64)
	c, _ := strconv.ParseFloat(input[2], 64)
	v0, _ := strconv.ParseFloat(input[3], 64)
	v1, _ := strconv.ParseFloat(input[4], 64)
	v2, _ := strconv.ParseFloat(input[5], 64)
	fmt.Printf("%.15f", minRoute(a, b, c, v0, v1, v2))
}
