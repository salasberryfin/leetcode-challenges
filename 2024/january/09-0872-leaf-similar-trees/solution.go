package main

import (
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
	}
	traverse(root.Left, leaves)
	traverse(root.Right, leaves)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1, leaves2 := make([]int, 0), make([]int, 0)
	traverse(root1, &leaves1)
	traverse(root2, &leaves2)

	return reflect.DeepEqual(leaves1, leaves2)
}
