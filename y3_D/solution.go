package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minCourse(c [][2]int, p float64) string {
	sort.Slice(c, func(i, j int) bool {
		return c[i][0] < c[j][0]
	})

	res := [2]int{}
	minVal := math.MaxFloat64
	i := len(c) - 1
	for i > 0 {
		ci := float64(c[i][0])
		t := ci / p

		l, r := 0, i-1
		for l < r {
			mid := (l + r) / 2
			cj := float64(c[mid][0])
			if cj < t {
				l = mid + 1
			} else {
				r = mid
			}
		}

		for k := l - 1; k < l+2; k++ {
			if k < 0 || k == i || k >= len(c) {
				continue
			}

			currVal := math.Abs((ci / float64(c[k][0])) - p)
			if currVal < minVal {
				res = [2]int{c[i][1], c[k][1]}
				minVal = currVal
			}
		}

		i--
	}

	return strconv.Itoa(res[0]+1) + " " + strconv.Itoa(res[1]+1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	np := strings.Split(strings.TrimSpace(line1), " ")
	n, _ := strconv.Atoi(np[0])
	p, _ := strconv.ParseFloat(np[1], 64)

	c := make([][2]int, n)
	line2, _ := reader.ReadString('\n')
	input := strings.Split(strings.TrimSpace(line2), " ")
	for i := range n {
		val, _ := strconv.Atoi(input[i])
		c[i] = [2]int{val, i}
	}

	writer.WriteString(minCourse(c, p))
}
