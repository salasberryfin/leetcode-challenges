package main

import "fmt"

// TreeNode is the definition of a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(tree *TreeNode) {
	// print nodes of the binary tree
	fmt.Println("Tree:", tree)
	if tree.Left != nil {
		preorderTraversal(tree.Left)
	}

	if tree.Right != nil {
		preorderTraversal(tree.Right)
	}
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// if one of the nodes is null, get the non-null one
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	// update value for current node
	root1.Val = root1.Val + root2.Val

	// get new left/right node
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}

func main() {
	result := mergeTrees(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		&TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
	)
	preorderTraversal(result)
}
