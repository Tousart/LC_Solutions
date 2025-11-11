package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxInterval(trios [][]float64) float64 {
	sort.Slice(trios, func(i, j int) bool {
		return trios[i][1] < trios[j][1]
	})

	weights := make([]float64, len(trios))
	weights[0] = trios[0][2]

	for i := 1; i < len(trios); i++ {
		leftBoard := trios[i][0]
		weight := trios[i][2]

		l, r := 0, i
		for l < r {
			mid := (l + r) / 2
			rightBoard := trios[mid][1]

			if rightBoard > leftBoard {
				r = mid
			} else if rightBoard <= leftBoard {
				l = mid + 1
			}
		}

		if l > 0 {
			weights[i] = max(weights[i-1], weight+weights[l-1])
		} else {
			weights[i] = max(weights[i-1], weight)
		}
	}

	return weights[len(weights)-1]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line1))

	if n == 0 {
		writer.WriteString("0")
		return
	}

	trios := make([][]float64, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		lrw := strings.Split(strings.TrimSpace(line), " ")
		l, _ := strconv.ParseFloat(lrw[0], 64)
		r, _ := strconv.ParseFloat(lrw[1], 64)
		w, _ := strconv.ParseFloat(lrw[2], 64)
		trios[i] = []float64{l, r, w}
	}

	writer.WriteString(strconv.FormatFloat(maxInterval(trios), 'f', 4, 64))
}
