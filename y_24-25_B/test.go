package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var result int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	q := make([]int, n)
	c := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		q[i], _ = strconv.Atoi(scanner.Text())
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		c[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	if a == b {
		for i := 0; i < n; i++ {
			result += q[i] * a
		}
	} else {
		for i := 0; i < n; i++ {
			val := float64(q[i]) * (((float64(c[i]) * float64((b - a))) / 255.0) + float64(a))
			result += int(math.Ceil(val))
		}
	}

	fmt.Print(result)
}
