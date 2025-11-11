package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func buyCloth(l, sum int, stores [][4]int) (int, []int) {
	maxVal := 10000001
	n := len(stores)

	// dp[i][j] = минимальная стоимость покупки j метров в [0; i] магазинах
	dp := make([][]int, n+1)
	prev := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
		prev[i] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			dp[i][j] = maxVal
			prev[i][j] = -1
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		p := stores[i-1][0]
		r := stores[i-1][1]
		q := stores[i-1][2]
		f := stores[i-1][3]

		for j := 0; j <= sum; j++ {
			if dp[i-1][j] == maxVal {
				continue
			}

			// покупаем k-метров в текущем магазине
			for k := 0; k <= f; k++ {
				if j+k > sum {
					break
				}
				pricePer := p
				if k >= r {
					pricePer = q
				}
				cost := dp[i-1][j] + k*pricePer
				if cost < dp[i][j+k] {
					dp[i][j+k] = cost
					prev[i][j+k] = k
				}
			}
		}
	}

	res := maxVal
	idx := -1
	for j := l; j <= sum; j++ {
		if dp[n][j] < res {
			res = dp[n][j]
			idx = j
		}
	}

	resSlice := make([]int, n)
	currIdx := idx
	for i := n; i >= 1; i-- {
		k := prev[i][currIdx]
		resSlice[i-1] = k
		currIdx -= k
	}

	return res, resSlice
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	nl := strings.Split(strings.TrimSpace(line1), " ")
	n, _ := strconv.Atoi(nl[0])
	l, _ := strconv.Atoi(nl[1])

	sum := 0
	idxPrice := make(map[int][2]int)
	stores := make([][4]int, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		prqf := strings.Split(strings.TrimSpace(line), " ")
		p, _ := strconv.Atoi(prqf[0])
		r, _ := strconv.Atoi(prqf[1])
		q, _ := strconv.Atoi(prqf[2])
		f, _ := strconv.Atoi(prqf[3])

		stores[i] = [4]int{p, r, q, f}
		idxPrice[i] = [2]int{0, 0}
		sum += f
	}

	if l == 0 {
		writer.WriteString("0")
		writer.WriteByte('\n')
		output := strings.Repeat("0 ", n)
		writer.WriteString(output[:len(output)-1])
		return
	}

	if sum < l {
		writer.WriteString("-1")
		return
	}

	res, resSlice := buyCloth(l, sum, stores)
	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
	for i := range resSlice {
		if i > 0 {
			writer.WriteString(" ")
		}
		writer.WriteString(strconv.Itoa(resSlice[i]))
	}
}
