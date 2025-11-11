package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func findPairs(coords, dxy map[[2]int]struct{}) int {
	var res int
	for coord := range coords {
		x, y := coord[0], coord[1]
		for shift := range dxy {
			if _, ok := coords[[2]int{x + shift[0], y + shift[1]}]; ok {
				res++
			}
		}
	}
	return res / 2
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	nd := strings.Split(strings.TrimSpace(line1), " ")
	n, _ := strconv.Atoi(nd[0])
	d, _ := strconv.Atoi(nd[1])

	coords := make(map[[2]int]struct{})
	for range n {
		line, _ := reader.ReadString('\n')
		xy := strings.Split(strings.TrimSpace(line), " ")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		coords[[2]int{x, y}] = struct{}{}
	}

	dxy := make(map[[2]int]struct{})
	shifts := [2]int{1, -1}
	for dx := 0; dx*dx <= d; dx++ {
		dyFloat := math.Sqrt(float64(d - dx*dx))
		if math.Trunc(dyFloat) == dyFloat {
			dy := int(dyFloat)
			for _, shiftX := range shifts {
				for _, shiftY := range shifts {
					dxy[[2]int{dx * shiftX, dy * shiftY}] = struct{}{}
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(findPairs(coords, dxy)))
}
