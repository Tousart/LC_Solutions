package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	node := &TreeNode{Val: nums[mid]}

	node.Left = buildTree(nums, left, mid-1)
	node.Right = buildTree(nums, mid+1, right)

	return node
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(nums))
}
