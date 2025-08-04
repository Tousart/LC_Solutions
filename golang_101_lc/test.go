package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return cmp(root.Left, root.Right)
}

func cmp(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	if !cmp(left.Left, right.Right) {
		return false
	} else if !cmp(left.Right, right.Left) {
		return false
	}

	return true
}

func main() {
	var root TreeNode
	root.Val = 1

	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}

	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}

	fmt.Println(isSymmetric(&root))
}
