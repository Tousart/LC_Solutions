package main

import (
	"fmt"
)

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	l := sortArray(nums[:mid])
	r := sortArray(nums[mid:])

	return merge(l, r)
}

func merge(l, r []int) []int {
	i, j := 0, 0
	res := make([]int, len(l)+len(r))
	idx := 0

	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			res[idx] = l[i]
			i++
		} else {
			res[idx] = r[j]
			j++
		}
		idx++
	}

	for i < len(l) {
		res[idx] = l[i]
		i++
		idx++
	}

	for j < len(r) {
		res[idx] = r[j]
		j++
		idx++
	}

	return res
}

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
}

// type MinHeap []int

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MinHeap) Push(x any) {
// 	*h = append(*h, x.(int))
// }

// func (h *MinHeap) Pop() any {
// 	x := (*h)[len(*h)-1]
// 	*h = (*h)[:len(*h)-1]
// 	return x
// }

// func sortArray(nums []int) []int {
// 	h := &MinHeap{}
// 	heap.Init(h)

// 	for _, num := range nums {
// 		heap.Push(h, num)
// 	}

// 	for i := range nums {
// 		nums[i] = heap.Pop(h).(int)
// 	}

// 	return nums
// }
