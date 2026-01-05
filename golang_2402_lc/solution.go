package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Meeting struct {
	roomNumber int
	endTime    int
}

type MinHeapMeetings []Meeting

func (h MinHeapMeetings) Len() int { return len(h) }
func (h MinHeapMeetings) Less(i, j int) bool {
	return (h[i].endTime < h[j].endTime) || (h[i].endTime == h[j].endTime && h[i].roomNumber < h[j].roomNumber)
}
func (h MinHeapMeetings) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeapMeetings) Push(x any) {
	*h = append(*h, x.(Meeting))
}

func (h *MinHeapMeetings) Pop() any {
	x := (*h)[len((*h))-1]
	*h = (*h)[:len((*h))-1]
	return x
}

type MinHeapRooms []int

func (h MinHeapRooms) Len() int { return len(h) }
func (h MinHeapRooms) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeapRooms) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeapRooms) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeapRooms) Pop() any {
	x := (*h)[len((*h))-1]
	*h = (*h)[:len((*h))-1]
	return x
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	rooms := MinHeapRooms{}
	heap.Init(&rooms)
	for i := range n {
		heap.Push(&rooms, i)
	}

	h := MinHeapMeetings{}
	heap.Init(&h)
	roomCount := make([]int, n)

	for _, m := range meetings {
		start, end := m[0], m[1]

		for h.Len() > 0 && h[0].endTime <= start {
			heap.Push(&rooms, h[0].roomNumber)
			// fmt.Printf("номер комнаты: %d, время выхода: %d\n", h[0].roomNumber, h[0].endTime)
			heap.Pop(&h)
		}

		var (
			room   int
			newEnd int
		)

		if rooms.Len() > 0 {
			room = heap.Pop(&rooms).(int)
			newEnd = end
		} else {
			r := heap.Pop(&h).(Meeting)
			room = r.roomNumber
			newEnd = r.endTime + (end - start)
		}
		roomCount[room]++
		heap.Push(&h, Meeting{roomNumber: room, endTime: newEnd})
	}

	var (
		res      int
		resCount int
	)

	// fmt.Println(roomCount)

	for number, count := range roomCount {
		if count > resCount {
			res = number
			resCount = count
		}
	}

	return res
}

func main() {
	n := 3
	meetings := [][]int{
		{1, 20},
		{2, 10},
		{3, 5},
		{4, 9},
		{6, 8},
	}
	fmt.Println(mostBooked(n, meetings))

	n1 := 2
	meetings1 := [][]int{
		{0, 10},
		{1, 2},
		{12, 14},
		{13, 15},
	}
	fmt.Println(mostBooked(n1, meetings1))

	n2 := 3
	meetings2 := [][]int{
		{0, 10},
		{1, 9},
		{2, 8},
		{3, 7},
		{4, 6},
	}
	fmt.Println(mostBooked(n2, meetings2))

	n3 := 5
	meetings3 := [][]int{
		{10, 14},
		{13, 17},
		{3, 12},
		{6, 13},
	}
	fmt.Println(mostBooked(n3, meetings3))
}
