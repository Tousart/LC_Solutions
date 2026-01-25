package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	index     int
	value     int
	sum       int
	left      *Node
	right     *Node
	heapIndex int
}

type MinHeap []*Node

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	if h[i].sum != h[j].sum {
		return h[i].sum < h[j].sum
	}
	return h[i].index < h[j].index
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIndex = i
	h[j].heapIndex = j
}

func (h *MinHeap) Push(x any) {
	n := len(*h)
	node := x.(*Node)
	node.heapIndex = n
	*h = append(*h, node)
}

func (h *MinHeap) Pop() any {
	node := (*h)[len(*h)-1]
	node.heapIndex = -1
	*h = (*h)[:len(*h)-1]
	return node
}

func minimumPairRemoval(nums []int) int {
	var prevNode *Node
	h := MinHeap{}
	heap.Init(&h)

	violations := 0
	for i, num := range nums {
		curr := &Node{
			index: i,
			value: num,
		}

		if prevNode != nil {
			prevNode.right = curr
			curr.left = prevNode

			if curr.value < prevNode.value {
				violations++
			}

			prevNode.sum = prevNode.value + curr.value
			heap.Push(&h, prevNode)
		}

		prevNode = curr
	}

	prevNode.sum = math.MaxInt64
	heap.Push(&h, prevNode)

	res := 0

	for violations > 0 {
		res++

		curr := heap.Pop(&h).(*Node)

		if curr.right == nil {
			continue
		}

		next := curr.right

		if next.value < curr.value {
			violations--
		}

		curr.value += next.value
		curr.right = next.right

		if next.right != nil {
			next.right.left = curr

			if next.right.value < next.value {
				violations--
			}

			if curr.value > next.right.value {
				violations++
			}

			curr.sum = curr.value + next.right.value
		} else {
			curr.sum = math.MaxInt64
		}

		if next.heapIndex >= 0 {
			heap.Remove(&h, next.heapIndex)
		}

		if curr.right != nil {
			heap.Push(&h, curr)
		}

		if curr.left != nil {
			prev := curr.left

			if prev.heapIndex >= 0 {
				heap.Remove(&h, prev.heapIndex)
			}

			if prev.value > curr.value-next.value {
				violations--
			}

			if prev.value > curr.value {
				violations++
			}

			prev.sum = prev.value + curr.value
			heap.Push(&h, prev)
		}
	}

	return res
}

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(minimumPairRemoval(nums))
}
