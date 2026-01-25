package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currA := headA
	currB := headB
	for currA != nil && currB != nil {
		currA = currA.Next
		currB = currB.Next
	}

	for currA != nil {
		headA = headA.Next
		currA = currA.Next
	}

	for currB != nil {
		headB = headB.Next
		currB = currB.Next
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func main() {
	headA := &ListNode{Val: 1}
	headA.Next = &ListNode{Val: 2}
	headB := &ListNode{Val: 2}
	fmt.Println(getIntersectionNode(headA, headB))
}

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	currA := headA
// 	currB := headB
// 	for currA != currB {
// 		if currA == nil {
// 			currA = headB
// 		} else {
// 			currA = currA.Next
// 		}

// 		if currB == nil {
// 			currB = headA
// 		} else {
// 			currB = currB.Next
// 		}
// 	}

// 	return currA
// }
