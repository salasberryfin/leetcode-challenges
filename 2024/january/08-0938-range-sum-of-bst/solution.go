package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, sum *int, l, h int) {
	if root == nil {
		return
	}
	if l <= root.Val && root.Val <= h {
		*sum += root.Val
	}
	traverse(root.Left, sum, l, h)
	traverse(root.Right, sum, l, h)
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var result int

	traverse(root, &result, low, high)

	return result
}
