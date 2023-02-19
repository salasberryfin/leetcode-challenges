/*
The strategy is very similar to the one followed to get the maximum depth of a tree
Starting from the maximum depth code:
	- checking both left and right for each node
	- returning 0 if the node is nil
Add a new variable that holds the maximum diameter so far
	- For each iteration check if `left+right` is larger than `max`
	- If it is, update max
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	var result func(*TreeNode) int
	result = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := result(root.Left)
		right := result(root.Right)
		if left+right > max {
			max = left + right
		}
		if left > right {
			return left + 1
		}

		return right + 1
	}

	result(root)

	return max
}
