package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	head := ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	fmt.Println(hasCycle(&head))
}

// func hasCycle(head *ListNode) bool {
// 	count := 0
// 	for head != nil {
// 		if count > 1e4 {
// 			return true
// 		}
// 		head = head.Next
// 		count++
// 	}
// 	return false
// }
