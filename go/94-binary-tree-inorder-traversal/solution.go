package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var traversal []int

	if root == nil {
		return traversal
	}

	traversal = append(traversal, inorderTraversal(root.Left)...)
	traversal = append(traversal, root.Val)
	traversal = append(traversal, inorderTraversal(root.Right)...)

	return traversal
}

func main() {
	root := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	result := inorderTraversal(&root)
	fmt.Printf("%v\n", result)
}
