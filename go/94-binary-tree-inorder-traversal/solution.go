/*
Inorder traversal: getting all elements from left side to right side
Recursive method:
	1. Check if root is nil: return
	2. Call recursively with tree.Left -> append result to slice of tree.Val
		a. Instead of printing, we're storing it in an array
		b. If printing -> print right after the recursive call with tree.Left
	3. Append current value to result slice -> this is what creates the left to right order
	4. Call recursively with tree.Right
*/
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
