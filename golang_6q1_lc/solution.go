package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (h intHeap) Len() int { return len(h) }

func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := intHeap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		a := heap.Pop(&h).(int)
		b := heap.Pop(&h).(int)

		if a != b {
			heap.Push(&h, a-b)
		}
	}

	if h.Len() == 1 {
		return h[0]
	}

	return 0
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))
}
