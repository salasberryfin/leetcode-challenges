/*
Use DFS to traverse the tree

- Start the DFS traverse for left and right nodes from the root.
- For each new node, check whether the zig zag path is longer going
left or right.
- Keep track of the longest path
- Check longest path between starting from root.Left and root.Right
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func visit(node *TreeNode, last string, count int) int {
	if node == nil {
		return count
	}

	fmt.Printf("Node %s: %d\n", last, node.Val)
	if last == "right" {
		// next move is left
		right := visit(node.Left, "left", count+1)
		left := visit(node.Right, "right", 0)
		if left > right {
			return left
		} else {
			return right
		}
	} else {
		// next move is right
		right := visit(node.Right, "right", count+1)
		left := visit(node.Left, "left", 0)
		if left > right {
			return left
		} else {
			return right
		}
	}
}

func longestZigZag(root *TreeNode) int {
	fmt.Println("Starting from right")
	right := visit(root.Right, "right", 0)
	fmt.Println("Starting from left")
	left := visit(root.Left, "left", 0)

	fmt.Println("Left max:", left, "Right max:", right)
	if left > right {
		return left
	}

	return right
}
