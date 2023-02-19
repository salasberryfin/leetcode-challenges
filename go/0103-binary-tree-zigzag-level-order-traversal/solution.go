/*
Level Order Traversal: going over each element at each level
To traverse the tree per level:
	- First get the height of the tree
	- Iterate over each level and create a function that gets the nodes in each
	- Add nodes of each level to the `levels` slice
	- To satisfy the zigzag: check whether the level is even or odd
		. Even levels go from left to right
		. Odd levels go from right to left
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := getHeight(node.Left)
	right := getHeight(node.Right)
	if left > right {
		return left + 1
	}

	return right + 1
}

func reverse(sl []int) []int {
	for i, j := 0, len(sl)-1; i <= j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}

	return sl
}

func getLevel(root *TreeNode, level int, nodes *[]int) {
	if root == nil {
		return
	}
	if level == 0 {
		*nodes = append(*nodes, root.Val)
	} else {
		getLevel(root.Left, level-1, nodes)
		getLevel(root.Right, level-1, nodes)
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	height := getHeight(root)
	nodes := [][]int{}
	for i := 0; i < height; i++ {
		level := &[]int{}
		getLevel(root, i, level)
		if i%2 != 0 {
			reverse(*level)
		}
		nodes = append(nodes, *level)
	}

	return nodes
}
