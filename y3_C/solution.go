package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func getIn(k, w, h float64, a, b []int, n int) bool {
	limW := w / k
	limH := h / k
	currentH := 0.0

	i := 0
	for i < n {
		currB := b[i]
		rows := 1
		sum := 0.0

		j := i
		for j < n && b[j] == currB {
			currA := float64(a[j])

			if currA > limW {
				return false
			}

			if currA+sum <= limW {
				sum += currA
			} else {
				sum = currA
				rows++
			}

			j++
		}

		currentH += float64(currB) * float64(rows)
		if currentH > limH {
			return false
		}

		i = j
	}

	return currentH <= limH
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	nwh := strings.Split(strings.TrimSpace(line1), " ")
	n, _ := strconv.Atoi(nwh[0])
	wInt, _ := strconv.Atoi(nwh[1])
	hInt, _ := strconv.Atoi(nwh[2])

	a := make([]int, n)
	b := make([]int, n)
	minA := math.MaxInt64
	minB := math.MaxInt64
	for i := range n {
		line, _ := reader.ReadString('\n')
		ab := strings.Split(strings.TrimSpace(line), " ")
		aInp, _ := strconv.Atoi(ab[0])
		a[i] = aInp
		minA = min(minA, aInp)
		bInp, _ := strconv.Atoi(ab[1])
		b[i] = bInp
		minB = min(minB, bInp)
	}

	w := float64(wInt)
	h := float64(hInt)
	right := min(w/float64(minA), h/float64(minB))
	left := 0.0

	for range 80 {
		mid := (left + right) / 2.0
		if getIn(mid, w, h, a, b, n) {
			left = mid
		} else {
			right = mid
		}
	}

	writer.WriteString(strconv.FormatFloat(left, 'f', 15, 64))
}
