package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printElements(head *ListNode) {
	for {
		fmt.Println(head)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := head
	var length int
	for {
		length++
		if newHead.Next == nil {
			break
		}
		newHead = newHead.Next
	}

	if length == n {
		//fmt.Println("Returning empty")
		return head.Next
	}

	nth := length - n - 1

	newHead = head
	var i int
	for i < nth {
		newHead = newHead.Next
		i++
	}
	newHead.Next = newHead.Next.Next

	return head
}

func main() {
	first := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	second := ListNode{
		Val:  1,
		Next: nil,
	}

	third := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	inputs := []ListNode{
		first,
		second,
		third,
	}

	positions := []int{
		2,
		1,
		1,
	}

	var i int
	for i < len(inputs) {
		fmt.Println("Calculating result for", &inputs[i], "and n =", positions[i])
		result := removeNthFromEnd(&inputs[i], positions[i])
		fmt.Println("Result:", result)
		i++
	}
}
