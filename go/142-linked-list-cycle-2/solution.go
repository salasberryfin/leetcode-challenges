/*
Check corner cases:
	1. Linked list is nil -> return nil
	2. Linked list does not cycle: find a node with node.Next == nil -> return nil
	3. A single node list with node.Next = node -> return node

Store each visited element in an slice of ListNode
	1. If node is in the visited list: this is the cycle.
	2. Break and return value of this node.
*/
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

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var posNode *ListNode

	listMap := map[*ListNode]int{}
	linkedList := head
	var i int
	for {
		if linkedList.Next == nil {
			// there's no cycle
			return nil
		} else if linkedList == linkedList.Next {
			// there's one node that cycles back to itself
			return linkedList
		}
		_, ok := listMap[linkedList.Next]
		_, currOk := listMap[linkedList]
		if ok && !currOk {
			// next node has already been visited but current has not
			// else we're going over the list again
			posNode = linkedList.Next
			break
		}
		listMap[linkedList] = i
		linkedList = linkedList.Next
		i++
	}

	return posNode
}
