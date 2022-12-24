package main

import (
	"fmt"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createStr(tree *TreeNode) string {
	var result string

	fmt.Printf("Current tree val: %d\n", tree.Val)

	result += "("
	result += fmt.Sprint(tree.Val)

	if tree.Right != nil {
		result += createStr(tree.Right)
	}
	if tree.Left != nil {
		if tree.Right == nil {
			result += "()"
		}
		result += createStr(tree.Left)
	}

	result += ")"

	return result
}

func tree2str(root *TreeNode) string {
	var result string
	result = createStr(root)
	result = strings.TrimPrefix(result, "(")
	result = strings.TrimSuffix(result, ")")

	return result
}

func main() {
	input1 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Printf("Result for input 1: %v\n", tree2str(&input1))

	input2 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	fmt.Printf("Result for input 2: %v\n", tree2str(&input2))
}
