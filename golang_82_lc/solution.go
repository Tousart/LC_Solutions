package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	prev := &ListNode{Val: 101}
	prev.Next = head
	currPrev := prev

	for currPrev.Next != nil && currPrev.Next.Next != nil {
		if currPrev.Next.Next.Val == currPrev.Next.Val {
			curr := currPrev.Next
			rmVal := curr.Val
			for curr != nil && curr.Val == rmVal {
				curr = curr.Next
			}
			currPrev.Next = curr
		} else {
			currPrev = currPrev.Next
		}
	}

	return prev.Next
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	fmt.Println(deleteDuplicates(head))

	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 1}
	fmt.Println(deleteDuplicates(head1))
}
