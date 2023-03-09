package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	headMap := map[*ListNode]bool{}
	current := head
	for {
		if current.Next == nil {
			// no cycle
			return nil
		} else if _, ok := headMap[current]; ok {
			// found cycle
			return current
		}
		headMap[current] = true
		current = current.Next
	}
}
