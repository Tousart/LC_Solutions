package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{nil}
	}

	return generate(1, n)
}

func generate(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}

	res := make([]*TreeNode, 0)

	for root := l; root <= r; root++ {
		left := generate(l, root-1)
		right := generate(root+1, r)
		for _, combL := range left {
			for _, combR := range right {
				res = append(res, &TreeNode{Val: root, Left: combL, Right: combR})
			}
		}
	}

	return res
}

func main() {
	n := 2
	fmt.Println(generateTrees(n))
}
