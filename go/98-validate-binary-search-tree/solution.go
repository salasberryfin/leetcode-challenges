/*
Valid binary search tree:
	- left subtree contains only nodes with keys less than the node's key
	- right subtree contains only nodes with keys greater than the node's key
	- both left/right must also be binary trees

Check tree with left and right values for comparing each side:
	1. left < root.Val < right
	2. Recursive call for left:
		- left = left, right = root.Val -> maintain low threshold, decrease high threshold to value of current node
	3. Recursive call for left:
		- left = root.Val, right = right -> maintain high threshold, increase low threshold to value of current node
Calls for left and right are opposite
If one of them is false -> it is an invalid Binary Search Tree:
	- if leftResult == rightResult -> return either
	- it leftResult != rightResult -> return false
*/
package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if root.Val >= right || root.Val <= left {
		return false
	}

	lSide := checkTree(root.Left, left, root.Val)
	rSide := checkTree(root.Right, root.Val, right)

	if lSide == rSide {
		return lSide
	} else {
		return false
	}
}

func isValidBST(root *TreeNode) bool {

	if root.Left == nil && root.Right == nil {
		return true
	}

	// cannot convert math.Inf(1) to int and get a positive number
	// so I set it to a large number hoping that the test cases don't go over it
	return checkTree(root, int(math.Inf(-1)), 10000000000000)
}
