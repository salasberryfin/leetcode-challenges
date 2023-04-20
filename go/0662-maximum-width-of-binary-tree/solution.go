package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node *TreeNode
	val  int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func widthOfBinaryTree(root *TreeNode) int {
	q := []Node{}
	q = append(q, Node{root, 1})

	maxWidth := 0
	for len(q) > 0 {
		// pop from queue
		left := q[0].val
		right := q[len(q)-1].val
		maxWidth = max(maxWidth, right-left+1)

		next := []Node{}
		// append new nodes to queue
		for _, pair := range q {
			node, val := pair.node, pair.val
			if node.Left != nil {
				next = append(next, Node{node.Left, (val * 2) - 1})
			}
			if node.Right != nil {
				next = append(next, Node{node.Right, val * 2})
			}
		}
		q = next
	}

	return maxWidth
}
