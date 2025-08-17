package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	next := dummy
	temp := dummy

	for next.Next != nil {
		if n <= 0 {
			temp = temp.Next
		}
		next = next.Next
		n--
	}

	temp.Next = temp.Next.Next

	return dummy.Next
}

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	next := removeNthFromEnd(&head, 2)
	for next != nil {
		fmt.Printf("%d ", next.Val)
		next = next.Next
	}
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	slice := []*ListNode{}
// 	next := head

// 	for next != nil {
// 		slice = append(slice, next)
// 		next = next.Next
// 	}

// 	ind := len(slice) - n - 1

// 	if ind < 0 {
// 		return head.Next
// 	}

// 	temp := slice[ind]
// 	temp.Next = temp.Next.Next
// 	return head
// }
