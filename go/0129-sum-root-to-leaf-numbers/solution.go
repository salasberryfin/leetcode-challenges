/*
*
DFS to traverse the binary tree: preorder traversal
To generate the number appending each new node: num*10 + node.Val

	e.g. 4 -> 5 -> 1: 4, 4*10 + 5 = 45, 45*10 + 1 = 451

Sum numbers on the left and right side
Check that we don't get to a nil, nil case which would return 0

*
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	val = val*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return val
	}

	return traverse(root.Left, val) + traverse(root.Right, val)
}

func sumNumbers(root *TreeNode) int {
	var result int
	result = traverse(root, 0)

	return result
}
