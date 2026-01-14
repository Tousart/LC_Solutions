package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)

	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range freqs {
		buckets[freq] = append(buckets[freq], num)
	}

	i := len(buckets) - 1
	for len(res) < k {
		if len(buckets[i]) > 0 {
			res = append(res, buckets[i]...)
		}
		i--
	}

	return res
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
