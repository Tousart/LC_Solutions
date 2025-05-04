package main

import "fmt"

type RecentCounter struct {
	requests []int
}

func Constructor() *RecentCounter {
	return &RecentCounter{requests: make([]int, 0)}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.requests = append(rc.requests, t)
	for rc.requests[0] < (t - 3000) {
		rc.requests = rc.requests[1:]
	}
	return len(rc.requests)
}

func main() {
	count := Constructor()
	fmt.Println(count.Ping(1))
	fmt.Println(count.Ping(100))
	fmt.Println(count.Ping(3001))
	fmt.Println(count.Ping(3002))
}
