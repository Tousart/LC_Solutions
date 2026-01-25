package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var (
		prev *TreeNode
		dfs  func(node *TreeNode)
	)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		dfs(node.Left)

		node.Right = prev
		node.Left = nil
		prev = node
	}

	dfs(root)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	fmt.Println(root)
	flatten(root)
	fmt.Println(root)
}
