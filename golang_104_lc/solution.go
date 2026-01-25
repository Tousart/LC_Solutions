package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	depth++
	return max(dfs(node.Left, depth), dfs(node.Right, depth))
}

func main() {
	root := &TreeNode{Val: 1}
	fmt.Println(maxDepth(root))
}
