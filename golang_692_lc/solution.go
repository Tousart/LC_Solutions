package main

import (
	"container/heap"
	"fmt"
)

type WordFreq struct {
	Word string
	Freq int
}

type MinHeap []*WordFreq

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	if h[i].Freq == h[j].Freq {
		return h[i].Word > h[j].Word
	}
	return h[i].Freq < h[j].Freq
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*WordFreq))
}

func (h *MinHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	h := MinHeap{}
	heap.Init(&h)
	for word, freq := range wordFreq {
		obj := &WordFreq{Word: word, Freq: freq}

		if h.Len() < k {
			heap.Push(&h, obj)
		} else {
			if (h[0].Freq < freq) ||
				(h[0].Freq == freq && h[0].Word > word) {
				heap.Pop(&h)
				heap.Push(&h, obj)
			}
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		obj := heap.Pop(&h).(*WordFreq)
		res[i] = obj.Word
	}

	return res
}

func main() {
	words := []string{"a", "b", "b"}
	k := 1
	fmt.Println(topKFrequent(words, k))
}
