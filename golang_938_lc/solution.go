package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var res int
	sum(root, low, high, &res)
	return res
}

func sum(root *TreeNode, low, high int, res *int) {
	if root == nil {
		return
	}

	if root.Val < low {
		sum(root.Right, low, high, res)
	} else if root.Val > high {
		sum(root.Left, low, high, res)
	} else {
		*res += root.Val
		sum(root.Right, low, high, res)
		sum(root.Left, low, high, res)
	}
}

func main() {
	root := TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 18}

	fmt.Println(rangeSumBST(&root, 7, 15))
}
