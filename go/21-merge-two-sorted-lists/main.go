package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinkedList(list *ListNode) {
	for {
		fmt.Printf("%d ", list.Val)
		if list.Next == nil {
			break
		}
		list = list.Next
	}
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	var result ListNode
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		result.Val = list1.Val
		result.Next = mergeTwoLists(list1.Next, list2)
	} else {
		result.Val = list2.Val
		result.Next = mergeTwoLists(list1, list2.Next)
	}

	return &result
}

func main() {
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	merged := mergeTwoLists(&list1, &list2)
	printLinkedList(merged)
}
