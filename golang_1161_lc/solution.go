package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	var (
		maxSum int = math.MinInt64
		level  int
	)

	// bfs
	queue := []*TreeNode{root}
	currLevel := 1
	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if sum > maxSum {
			maxSum = sum
			level = currLevel
		}

		currLevel++
	}

	return level
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: -8}
	root.Right = &TreeNode{Val: 0}
	fmt.Println(maxLevelSum(&root))
}
