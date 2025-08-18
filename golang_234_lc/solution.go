package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	step, temp := head, head
	for step.Next != nil && step.Next.Next != nil {
		temp = temp.Next
		step = step.Next.Next
	}

	temp = reverse(temp.Next)

	for temp != nil {
		if temp.Val != head.Val {
			return false
		}
		temp = temp.Next
		head = head.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 1}
	fmt.Println(isPalindrome(&head))
}
