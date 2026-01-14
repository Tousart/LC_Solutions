package main

import (
	"fmt"
)

func maximalRectangle(matrix [][]byte) int {
	var res int

	heights := make([]int, len(matrix[0]))
	for row := range matrix {
		for col := range matrix[0] {
			if matrix[row][col] == '1' {
				heights[col]++
			} else {
				heights[col] = 0
			}
		}

		res = max(res, currHighMaxArea(heights))
	}

	return res
}

func currHighMaxArea(heights []int) int {
	var area int
	stack := make([]int, 0)

	for i := range heights {
		for (len(stack) > 0) && (heights[i] < heights[stack[len(stack)-1]]) {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				area = max(area, (i-stack[len(stack)-1]-1)*h)
				fmt.Println(i, stack[len(stack)-1])
			} else {
				area = max(area, i*h)
			}
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		h := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]

		if len(stack) > 0 {
			area = max(area, h*(len(heights)-stack[len(stack)-1]-1))
		} else {
			area = max(area, h*len(heights))
		}
	}

	return area
}

func main() {
	matrix := [][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10010"),
	}
	fmt.Println(maximalRectangle(matrix))
}

// func maximalRectangle(matrix [][]byte) int {
// 	var res int

// 	rec := make([][]int, len(matrix))
// 	for row := range len(matrix) {
// 		rec[row] = make([]int, len(matrix[0]))
// 		rec[row][0] = int(matrix[row][0] - '0')

// 		for col := 1; col < len(matrix[0]); col++ {
// 			rec[row][col] = int(matrix[row][col] - '0')
// 			if matrix[row][col] == '1' {
// 				rec[row][col] += rec[row][col-1]
// 			}
// 		}
// 	}

// 	for row := range len(matrix) {
// 		for col := range len(matrix[0]) {
// 			if matrix[row][col] == '1' {
// 				width := rec[row][col]
// 				res = max(res, width)

// 				i := row + 1
// 				for i < len(matrix) {
// 					if matrix[i][col] != '1' {
// 						break
// 					}

// 					width = min(width, rec[i][col])
// 					res = max(res, width*(i-row+1))

// 					i++
// 				}
// 			}
// 		}
// 	}

// 	return res
// }
