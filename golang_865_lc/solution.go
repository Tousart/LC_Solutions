package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	resNode := &TreeNode{}
	maxDeep := 0
	dfs(root, &maxDeep, 0, resNode)
	return resNode
}

func dfs(node *TreeNode, maxDeep *int, deep int, resNode *TreeNode) int {
	if node == nil {
		return deep - 1
	}

	*maxDeep = max(*maxDeep, deep)

	l := dfs(node.Left, maxDeep, deep+1, resNode)
	r := dfs(node.Right, maxDeep, deep+1, resNode)
	if l == *maxDeep && r == *maxDeep {
		*resNode = *node
	}

	return max(l, r)
}

func main() {
	root := &TreeNode{Val: 1}
	fmt.Println(subtreeWithAllDeepest(root).Val)
}
