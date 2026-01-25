package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)

		for i := range n {
			if queue[i] != nil {
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
				queue[i].Left, queue[i].Right = queue[i].Right, queue[i].Left
			}
		}

		queue = queue[n:]
	}

	return root
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 1}
	fmt.Println(invertTree(root))
}
