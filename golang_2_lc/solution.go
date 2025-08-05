package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		shift  int
		result ListNode
	)
	next := &result

	for l1 != nil || l2 != nil || shift != 0 {
		sum := shift

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		shift = sum / 10
		next.Next = &ListNode{Val: sum % 10}
		next = next.Next
	}

	return result.Next
}

func main() {
	l1 := ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	result := addTwoNumbers(&l1, &l2)

	next := result
	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
}
