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

	result += "("
	result += fmt.Sprint(tree.Val)

	if tree.Left != nil {
		result += createStr(tree.Left)
	}
	if tree.Right != nil {
		if tree.Left == nil {
			result += "()"
		}
		result += createStr(tree.Right)
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
		Right: &TreeNode{
			Val:   3,
			Right: nil,
			Left:  nil,
		},
		Left: &TreeNode{
			Val:   2,
			Right: nil,
			Left: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
		},
	}

	fmt.Printf("Result for input 1: %v\n", tree2str(&input1))

	input2 := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Right: nil,
			Left:  nil,
		},
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	}

	fmt.Printf("Result for input 2: %v\n", tree2str(&input2))
}
