package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := ListNode{}
	next := &result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			next.Next = list1
			list1 = list1.Next
		} else {
			next.Next = list2
			list2 = list2.Next
		}
		next = next.Next
	}

	if list1 != nil {
		next.Next = list1
	}

	if list2 != nil {
		next.Next = list2
	}

	return result.Next
}

func main() {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	next := mergeTwoLists(&l1, &l2)
	for next != nil {
		fmt.Printf("%d ", next.Val)
		next = next.Next
	}
}
