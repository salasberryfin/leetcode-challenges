package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversalBalanced(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	lDepth, leftBalanced := preOrderTraversalBalanced(node.Left)
	rDepth, rightBalanced := preOrderTraversalBalanced(node.Right)

	if !leftBalanced || !rightBalanced {
		return 0, false
	}

	if abs(lDepth, rDepth) > 1 {
		return 0, false
	}

	return max(lDepth, rDepth) + 1, true
}

func max(left, right int) int {
	if left > right {
		return left
	}

	return right
}

func abs(left, right int) int {
	if left > right {
		return left - right
	}

	return right - left
}

func isBalanced(root *TreeNode) bool {
	_, balanced := preOrderTraversalBalanced(root)

	return balanced
}
