package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	newHead := head
	var length int
	var locations []*ListNode
	for {
		length++
		locations = append(locations, newHead)
		if newHead.Next == nil {
			break
		}
		newHead = newHead.Next
	}

	return locations[length/2]
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
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	inputs := []ListNode{
		first,
		second,
	}

	var i int
	for i < len(inputs) {
		result := middleNode(&inputs[i])
		fmt.Println("Calculating result for", &inputs[i])
		fmt.Println("Middle element is", result)
		i++
	}
}
