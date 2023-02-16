package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(tree *TreeNode) int {
	if tree == nil {
		return 0
	}

	depthLeft := dfs(tree.Left)
	depthRight := dfs(tree.Right)

	if depthLeft > depthRight {
		return depthLeft + 1
	} else {
		return depthRight + 1
	}
}

func maxDepth(root *TreeNode) int {
	result := dfs(root)

	return result
}
