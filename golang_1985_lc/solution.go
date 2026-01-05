package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []string

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return (len(h[i]) > len(h[j])) || (len(h[i]) == len(h[j]) && h[i] > h[j])
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(string))
}

func (h *MaxHeap) Pop() any {
	x := (*h)[len((*h))-1]
	*h = (*h)[:len((*h))-1]
	return x
}

func kthLargestNumber(nums []string, k int) string {
	h := MaxHeap{}
	h = append(h, nums...)
	heap.Init(&h)

	for i := 1; i < k; i++ {
		heap.Pop(&h)
	}

	return h[0]
}

func main() {
	nums := []string{"3", "6", "7", "10"}
	k := 4
	fmt.Println(kthLargestNumber(nums, k))

	nums1 := []string{"2", "21", "12", "1"}
	k1 := 3
	fmt.Println(kthLargestNumber(nums1, k1))

	nums2 := []string{"1", "0", "0"}
	k2 := 1
	fmt.Println(kthLargestNumber(nums2, k2))
}
