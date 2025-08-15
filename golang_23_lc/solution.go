package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	l1 := merge(lists[:mid])
	l2 := merge(lists[mid:])

	res := ListNode{Val: -1}
	curr := &res

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return res.Next
}

func main() {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 5}

	l2 := ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	l3 := ListNode{Val: 2}
	l3.Next = &ListNode{Val: 6}

	lists := []*ListNode{&l1, &l2, &l3}

	next := mergeKLists(lists)
	for next != nil {
		fmt.Printf("%d ", next.Val)
		next = next.Next
	}
}
