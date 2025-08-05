package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt, math.MaxInt)
}

func validate(root *TreeNode, leftBord, rightBord int) bool {
	if root == nil {
		return true
	}

	if root.Val <= leftBord || root.Val >= rightBord {
		return false
	}

	if !validate(root.Left, leftBord, root.Val) {
		return false
	} else if !validate(root.Right, root.Val, rightBord) {
		return false
	}

	return true
}

func main() {
	root := TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(isValidBST(&root))
}
