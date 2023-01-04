package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinkedList(list *ListNode) []int {
	result := []int{}
	for list != nil {
		result = append(result, list.Val)
		fmt.Printf("%d ", list.Val)
		if list.Next == nil {
			break
		}
		list = list.Next
	}

	return result
}

func middleNode(head *ListNode) *ListNode {
	var middle *ListNode

	var headCopy []*ListNode
	depth := 0
	linkedList := head
	for linkedList != nil {
		headCopy = append(headCopy, linkedList)
		linkedList = linkedList.Next
		depth++
	}
	middleIndex := depth / 2
	middle = headCopy[middleIndex]

	return middle
}
