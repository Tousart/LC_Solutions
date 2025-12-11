package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *[]int) {
	if node.Left != nil {
		dfs(node.Left, res)
	}

	(*res) = append((*res), node.Val)

	if node.Right != nil {
		dfs(node.Right, res)
	}
}

func main() {
	root := TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(inorderTraversal(&root))
}
