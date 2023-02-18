package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	inorderTraversal(node.Left)
	fmt.Println("Node ", node.Val)
	inorderTraversal(node.Right)
}

func postOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	postOrderTraversal(node.Left)
	postOrderTraversal(node.Right)
	fmt.Println("Node ", node.Val)
}

// traverse performs Preorder Traversal on the binary tree and switches left and right
// to invert the nodes of the tree
func traverse(node *TreeNode) {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left
	traverse(node.Left)
	traverse(node.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	//traverse(root)
	inorderTraversal(root)

	return root
}
