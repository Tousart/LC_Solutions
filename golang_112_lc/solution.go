package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if !hasPathSum(root.Left, targetSum) {
		return hasPathSum(root.Right, targetSum)
	}

	return true
}

func main() {
	root := &TreeNode{Val: -2}
	root.Right = &TreeNode{Val: -3}
	targetSum := -5
	fmt.Println(hasPathSum(root, targetSum))
}
