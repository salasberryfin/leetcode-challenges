package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return traverse(root.Left, min, root.Val) && traverse(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return traverse(root, -math.MaxInt, math.MaxInt)
}
