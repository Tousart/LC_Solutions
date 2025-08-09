package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	}
	return right
}

func main() {
	root := TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}

	p := root.Left
	q := root.Right

	res := lowestCommonAncestor(&root, p, q)
	fmt.Println(res.Val)
}
