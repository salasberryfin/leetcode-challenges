package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	inOrderTraversal(node.Left, values)
	*values = append(*values, node.Val)
	inOrderTraversal(node.Right, values)
}

func minDiffInBST(root *TreeNode) int {
	minDist := math.MaxInt

	values := []int{}
	inOrderTraversal(root, &values)

	for i := 0; i+1 < len(values); i++ {
		if values[i+1]-values[i] < minDist {
			minDist = values[i+1] - values[i]
		}
	}

	return minDist
}
