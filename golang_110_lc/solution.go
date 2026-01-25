package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res := true

	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth
		}

		depth++
		l := dfs(node.Left, depth)
		r := dfs(node.Right, depth)

		if math.Abs(float64(l-r)) > 1 {
			res = false
		}

		return max(l, r)
	}

	dfs(root, 0)

	return res
}

// func dfs(node *TreeNode, depth int, flag *bool) int {
// 	if node == nil {
// 		return depth
// 	}

// 	depth++
// 	l := dfs(node.Left, depth, flag)
// 	r := dfs(node.Right, depth, flag)

// 	if math.Abs(float64(l-r)) > 1 {
// 		*flag = false
// 	}

// 	return max(l, r)
// }

func main() {
	root := &TreeNode{Val: 228}
	fmt.Println(isBalanced(root))
}
