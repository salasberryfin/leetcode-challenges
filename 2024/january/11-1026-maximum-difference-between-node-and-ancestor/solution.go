package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxVal(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func minVal(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func traverse(node *TreeNode, min, max int) int {
	if node == nil {
		return max - min
	}

	min = minVal(min, node.Val)
	max = maxVal(max, node.Val)

	return (maxVal(traverse(node.Left, min, max), traverse(node.Right, min, max)))
}

func maxAncestorDiff(root *TreeNode) int {

	return traverse(root, root.Val, root.Val)
}
