package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
	node  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		stack: make([]*TreeNode, 0),
		node:  root,
	}
}

func (this *BSTIterator) Next() int {
	for this.node != nil {
		this.stack = append(this.stack, this.node)
		this.node = this.node.Left
	}

	x := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.node = x.Right

	return x.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0 || this.node != nil
}

func main() {
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 20}
	obj := Constructor(root)
	fmt.Println(obj.Next())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
}

// type BSTIterator struct {
// 	nodes []int
// }

// func Constructor(root *TreeNode) BSTIterator {
// 	nodes := make([]int, 0)

// 	var dfs func(node *TreeNode)
// 	dfs = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		dfs(node.Left)
// 		nodes = append(nodes, node.Val)
// 		dfs(node.Right)
// 	}

// 	dfs(root)

// 	return BSTIterator{
// 		nodes: nodes,
// 	}
// }

// func (this *BSTIterator) Next() int {
// 	x := this.nodes[0]
// 	this.nodes = this.nodes[1:]
// 	return x
// }

// func (this *BSTIterator) HasNext() bool {
// 	return len(this.nodes) > 0
// }
