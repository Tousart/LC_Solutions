package main

import "fmt"

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent, size []int, a, b int) {
	a = find(parent, a)
	b = find(parent, b)

	if a == b {
		return
	}

	if size[a] < size[b] {
		parent[a] = b
	} else {
		parent[b] = a

		if size[a] == size[b] {
			size[a]++
		}
	}
}

func latestDayToCross(row int, col int, cells [][]int) int {
	shifts := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	n := row * col
	top, bottom := n, n+1

	parent := make([]int, n+2)
	for i := range n + 2 {
		parent[i] = i
	}

	size := make([]int, n+2)
	matrix := make([]bool, n)

	for day := n - 1; day >= 0; day-- {
		r, c := cells[day][0]-1, cells[day][1]-1

		idx := r*col + c
		matrix[idx] = true

		if r == 0 {
			union(parent, size, idx, top)
		}

		if r == row-1 {
			union(parent, size, idx, bottom)
		}

		for shift := range 4 {
			sr := r + shifts[shift][0]
			sc := c + shifts[shift][1]
			if sr >= 0 && sr < row && sc >= 0 && sc < col {
				neighbourIdx := sr*col + sc
				if matrix[neighbourIdx] {
					union(parent, size, idx, neighbourIdx)
				}
			}
		}

		if find(parent, top) == find(parent, bottom) {
			return day
		}
	}

	return 1
}

func main() {
	row := 6
	col := 2
	cells := [][]int{
		{4, 2},
		{6, 2},
		{2, 1},
		{4, 1},
		{6, 1},
		{3, 1},
		{2, 2},
		{3, 2},
		{1, 1},
		{5, 1},
		{5, 2},
		{1, 2},
	}
	fmt.Println(latestDayToCross(row, col, cells))
}
