/*
Go through left side and right side part of the tree
For each left side node:
  - left == right
  - recursively call for (left.Left, right.Right) and (left.Right, right.Left) to check mirrored tree
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}
