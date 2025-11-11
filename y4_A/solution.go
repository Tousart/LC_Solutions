package main

import (
	"bufio"
	"container/heap"
	"os"
	"sort"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func toMinutes(s string) int {
	h, _ := strconv.Atoi(s[:2])
	m, _ := strconv.Atoi(s[3:])
	return h*60 + m
}

func busCount(first, second [][2]int) int {
	freeA, freeB := &IntHeap{}, &IntHeap{}
	heap.Init(freeA)
	heap.Init(freeB)
	i, j := 0, 0
	buses := 0

	for i < len(first) || j < len(second) {
		if j == len(second) || (i < len(first) && first[i][0] <= second[j][0]) {
			// Рейс из A в B
			if freeA.Len() > 0 && (*freeA)[0] <= first[i][0] {
				heap.Pop(freeA)
			} else {
				buses++
			}
			heap.Push(freeB, first[i][1])
			i++
		} else {
			// Рейс из B в A
			if freeB.Len() > 0 && (*freeB)[0] <= second[j][0] {
				heap.Pop(freeB)
			} else {
				buses++
			}
			heap.Push(freeA, second[j][1])
			j++
		}
	}

	return buses
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line1))
	first := make([][2]int, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		input := strings.Split(strings.TrimSpace(line), "-")
		first[i] = [2]int{toMinutes(input[0]), toMinutes(input[1])}
	}
	sort.Slice(first, func(i, j int) bool {
		return first[i][0] < first[j][0]
	})

	line2, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(line2))
	second := make([][2]int, m)
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		input := strings.Split(strings.TrimSpace(line), "-")
		second[i] = [2]int{toMinutes(input[0]), toMinutes(input[1])}
	}
	sort.Slice(second, func(i, j int) bool {
		return second[i][0] < second[j][0]
	})

	writer.WriteString(strconv.Itoa(busCount(first, second)))
}
