package main

import "fmt"

type Router struct {
	queue    [][3]int
	packets  map[string]struct{}
	destTime map[int][]int
	limit    int
}

func Constructor(memoryLimit int) Router {
	return Router{
		queue:    [][3]int{},
		packets:  make(map[string]struct{}),
		destTime: make(map[int][]int),
		limit:    memoryLimit,
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	packet := fmt.Sprintf("%d,%d,%d", source, destination, timestamp)
	if _, ok := this.packets[packet]; ok {
		return false
	}

	if len(this.queue) == this.limit {
		oldPacket := this.queue[0]
		this.queue = this.queue[1:]
		delete(this.packets, fmt.Sprintf("%d,%d,%d", oldPacket[0], oldPacket[1], oldPacket[2]))
		this.destTime[oldPacket[1]] = this.destTime[oldPacket[1]][1:]
	}

	this.queue = append(this.queue, [3]int{source, destination, timestamp})
	this.packets[packet] = struct{}{}
	this.destTime[destination] = append(this.destTime[destination], timestamp)

	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.queue) > 0 {
		s, d, t := this.queue[0][0], this.queue[0][1], this.queue[0][2]
		this.queue = this.queue[1:]
		delete(this.packets, fmt.Sprintf("%d,%d,%d", s, d, t))
		this.destTime[d] = this.destTime[d][1:]
		return []int{s, d, t}
	} else {
		return []int{}
	}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	times := this.destTime[destination]

	if len(times) == 0 || times[0] > endTime || times[len(times)-1] < startTime {
		return 0
	}

	left, right := 0, len(times)-1
	for left < right {
		mid := (right + left) / 2
		if times[mid] < startTime {
			left = mid + 1
		} else {
			right = mid
		}
	}

	ind1 := right
	if times[ind1] > endTime {
		return 0
	}

	left = right
	right = len(times) - 1
	for left < right {
		mid := (right + left + 1) / 2
		if times[mid] > endTime {
			right = mid - 1
		} else {
			left = mid
		}
	}

	ind2 := left

	return ind2 - ind1 + 1
}

func main() {
	memoryLimit := 3
	obj := Constructor(memoryLimit)
	fmt.Println(obj.AddPacket(1, 77, 1))
	fmt.Println(obj.AddPacket(1, 77, 2))
	fmt.Println(obj.AddPacket(1, 77, 5))
	fmt.Println(obj.ForwardPacket())
	fmt.Println(obj.GetCount(77, 1, 3))
}
