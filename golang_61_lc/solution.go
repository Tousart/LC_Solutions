package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	var (
		last   *ListNode
		length int = 1
	)

	curr := head
	for curr.Next != nil {
		length++
		curr = curr.Next
	}
	last = curr

	realRotate := k % length
	if realRotate == 0 {
		return head
	}

	lastResNodeIdx := length - realRotate - 1
	currNodeIdx := 0
	currNode := head
	for currNodeIdx < lastResNodeIdx {
		currNode = currNode.Next
		currNodeIdx++
	}

	res := currNode.Next
	(*currNode).Next = nil
	(*last).Next = head
	return res
}

func main() {
	l := ListNode{Val: 0}
	l.Next = &ListNode{Val: 1}
	l.Next.Next = &ListNode{Val: 2}
	k := 4

	res := rotateRight(&l, k)
	curr := res
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}

	fmt.Println()

	l1 := ListNode{Val: 0}
	l1.Next = &ListNode{Val: 1}
	k1 := 4

	res1 := rotateRight(&l1, k1)
	curr1 := res1
	for curr1 != nil {
		fmt.Printf("%d ", curr1.Val)
		curr1 = curr1.Next
	}
}
