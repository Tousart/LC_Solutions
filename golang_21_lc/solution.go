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
			next.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			next.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		next = next.Next
	}

	for list1 != nil {
		next.Next = &ListNode{Val: list1.Val}
		next = next.Next
		list1 = list1.Next
	}

	for list2 != nil {
		next.Next = &ListNode{Val: list2.Val}
		next = next.Next
		list2 = list2.Next
	}

	return result.Next
}

func main() {
	list1 := ListNode{Val: 1}
	list1.Next = &ListNode{Val: 4}
	list2 := ListNode{Val: 2}

	res := mergeTwoLists(&list1, &list2)
	next := res
	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
}
