package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) bool {
	if root == nil {
		return true
	}

	emptyFound := false
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == nil {
			emptyFound = true
		} else {
			if emptyFound {
				return false
			}
			queue = append(queue, current.Left, current.Right)
		}
	}

	return true
}

func isCompleteTree(root *TreeNode) bool {
	return bfs(root)
}
