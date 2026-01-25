package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}

	if p == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	return isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}

	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 3}

	fmt.Println(isSameTree(p, q))
}
