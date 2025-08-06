package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n     int
		m     int
		x     int
		y     int
		indNX int
		indMY int
		count int
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	x, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ = strconv.Atoi(scanner.Text())
	bord := (x*y + 1) / 2

	house := make([]string, n*x)
	for i := 0; i < n*x; i++ {
		scanner.Scan()
		house[i] = scanner.Text()
	}

	for indNX < n*x {
		for indMY < m*y {
			w := 0
			for k := indNX; k < indNX+x; k++ {
				for l := indMY; l < indMY+y; l++ {
					if house[k][l] == 'X' {
						w += 1
					}
				}
			}

			if w >= bord {
				count += 1
			}

			indMY += y
		}

		indMY = 0
		indNX += x
	}

	fmt.Print(count)
}
