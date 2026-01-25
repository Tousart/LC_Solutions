package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	p := root.Left
	q := root.Right
	fmt.Println(lowestCommonAncestor(root, p, q))
}
