package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var traversal = []int{}

	if root == nil {
		return traversal
	}

	traversal = append(traversal, root.Val)

	traversal = append(traversal, preorderTraversal(root.Left)...)
	traversal = append(traversal, preorderTraversal(root.Right)...)

	return traversal
}
