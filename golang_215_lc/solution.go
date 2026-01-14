package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := MinHeap{}
	heap.Init(&h)

	for _, num := range nums {
		if h.Len() < k {
			heap.Push(&h, num)
		} else {
			if h[0] < num {
				heap.Pop(&h)
				heap.Push(&h, num)
			}
		}
	}

	return h[0]
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}
