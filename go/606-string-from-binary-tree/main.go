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
