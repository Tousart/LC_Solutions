package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	stop := false
	depth := 0

	for len(queue) > 0 {
		n := len(queue)
		depth++

		for i := range n {
			if queue[i].Left == nil && queue[i].Right == nil {
				stop = true
				break
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		if stop {
			break
		}

		queue = queue[n:]
	}

	return depth
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	fmt.Println(minDepth(root))
}
