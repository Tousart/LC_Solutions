package main

import (
	"container/heap"
	"fmt"
)

type Set struct {
	sum int
	i   int
	j   int
}

type SetHeap []Set

func (h SetHeap) Len() int           { return len(h) }
func (h SetHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h SetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SetHeap) Push(x any) {
	*h = append(*h, x.(Set))
}

func (h *SetHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &SetHeap{}

	for i, num := range nums1 {
		*h = append(*h, Set{sum: num + nums2[0], i: i, j: 0})
	}

	heap.Init(h)

	res := make([][]int, 0)

	for k > 0 {
		curr := heap.Pop(h).(Set)
		i, j := curr.i, curr.j
		res = append(res, []int{nums1[i], nums2[j]})
		k--

		j++
		if j < len(nums2) {
			heap.Push(h, Set{sum: nums1[i] + nums2[j], i: i, j: j})
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}
