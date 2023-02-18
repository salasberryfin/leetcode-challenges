package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	var newHead *ListNode
	for head != nil {
		newHead = &ListNode{
			Val:  head.Val,
			Next: previous,
		}
		previous = newHead
		head = head.Next
	}

	return newHead
}
