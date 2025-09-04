package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	dfs(root, &res, []string{})
	return res
}

func dfs(root *TreeNode, res *[]string, route []string) {
	if root == nil {
		return
	}

	route = append(route, strconv.Itoa(root.Val))

	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(route, "->"))
	}

	dfs(root.Left, res, route)
	dfs(root.Right, res, route)
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(binaryTreePaths(&root))
}
