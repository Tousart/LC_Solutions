package main

import "fmt"

func trap(height []int) int {
	var (
		level int
		res   int
		l     int
		r     int = len(height) - 1
	)

	for (l + 1) < r {
		minHigh := min(height[l], height[r])
		if minHigh > level {
			res += (minHigh - level) * (r - l - 1)
			level = minHigh
		}

		if height[l] < height[r] {
			l++
			res -= min(level, height[l])
		} else {
			r--
			res -= min(level, height[r])
		}
	}

	return res
}

func main() {
	height := []int{4, 2, 3}
	fmt.Println(trap(height))
}
