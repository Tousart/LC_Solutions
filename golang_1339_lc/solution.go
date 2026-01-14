package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	var res int

	mod := 1000000007

	var dfsSum func(node *TreeNode) int
	dfsSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return node.Val + dfsSum(node.Left) + dfsSum(node.Right)
	}

	sum := dfsSum(root)

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := dfs(node.Left)
		r := dfs(node.Right)
		currNodeSum := node.Val + l + r
		res = max(res, (sum-currNodeSum)*currNodeSum)

		return currNodeSum
	}

	dfs(root)

	return res % mod
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	fmt.Println(maxProduct(&root))
}

// type LeftRightSum struct {
// 	LeftSum  int
// 	RightSum int
// }

// func maxProduct(root *TreeNode) int {
// 	var res int

// 	mod := 1000000007
// 	sums := make(map[*TreeNode]LeftRightSum)
// 	sum := dfs(root, sums)

// 	for _, lrSum := range sums {
// 		l, r := lrSum.LeftSum, lrSum.RightSum
// 		res = max(res, max(l*(sum-l), r*(sum-r)))
// 	}

// 	return res % mod
// }

// func dfs(root *TreeNode, sums map[*TreeNode]LeftRightSum) int {
// 	if root == nil {
// 		return 0
// 	}

// 	l := dfs(root.Left, sums)
// 	r := dfs(root.Right, sums)

// 	sums[root] = LeftRightSum{LeftSum: l, RightSum: r}

// 	return root.Val + l + r
// }
