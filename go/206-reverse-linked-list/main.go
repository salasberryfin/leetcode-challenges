package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinkedList(list *ListNode) []int {
	var result []int
	for {
		result = append(result, list.Val)
		if list.Next == nil {
			break
		}
		list = list.Next
	}

	fmt.Printf("Printing linked list: %v\n", result)

	return result
}

func reverse(values []int, pos int) *ListNode {
	var reversed ListNode

	reversed.Val = values[pos]
	if pos > 0 {
		reversed.Next = reverse(values, pos-1)
	}

	return &reversed
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	var values []int
	temp := head
	if head == nil {
		return result
	}
	for temp != nil {
		values = append(values, temp.Val)
		temp = temp.Next
	}

	result = reverse(values, len(values)-1)

	return result
}
