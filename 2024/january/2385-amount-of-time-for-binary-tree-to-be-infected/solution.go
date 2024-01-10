package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func traverse(root *TreeNode, start int, minutes *int) (int, bool) {
	if root == nil {
		return 0, false
	}

	depthLeft, inLeft := traverse(root.Left, start, minutes)
	depthRight, inRight := traverse(root.Right, start, minutes)

	if root.Val == start {
		*minutes = max(*minutes, max(depthLeft, depthRight))
		return 1, true
	}

	if inLeft {
		*minutes = max(*minutes, depthRight+depthLeft)
		return depthLeft + 1, true
	} else if inRight {
		*minutes = max(*minutes, depthLeft+depthRight)
		return depthRight + 1, true
	} else {
		return max(depthLeft, depthRight) + 1, false
	}

}

func amountOfTime(root *TreeNode, start int) int {
	var minutes int
	traverse(root, start, &minutes)

	return minutes
}
