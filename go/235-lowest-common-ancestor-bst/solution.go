/*
This is a tough one.

Lowes Common Ancestor (LCA): lowest node in the tree that has both p, q as children
To find the LCA:
	1. Find path to p and q.
		- Go over tree recursively looking for given descendant
		- Check if is descendant of root.Left or root.Right and call recursively for the correct side
		- Append results to slice and generate path slice: nodes to p and nodes to q
	2. Find LCA in paths (len of path slices may be different):
		- Loop over paths and find the first element that is not equal in both slices
		- Store all checked nodes in slice
		- If the loop goes over all elements and does not find a difference, we return the last element in the checked slice
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) []int {
	var result = []int{}
	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = append(result, printTree(root.Left)...)
	result = append(result, printTree(root.Right)...)

	return result
}

func isChild(root, candidate *TreeNode) bool {
	if root == nil {
		return false
	}

	if root == candidate {
		return true
	}

	if (root.Left != nil && isChild(root.Left, candidate)) || (root.Right != nil && isChild(root.Right, candidate)) {
		return true
	}

	return false
}

func findPath(root, descent *TreeNode) []TreeNode {
	var result = []TreeNode{}

	if root == nil {
		return result
	}

	result = append(result, *root)

	if root == descent {
		return result
	}
	if isChild(root.Left, descent) {
		result = append(result, findPath(root.Left, descent)...)
	} else {
		result = append(result, findPath(root.Right, descent)...)
	}

	return result
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root.Left == nil || root.Right == nil {
		return root
	}

	pPath := findPath(root, p)
	qPath := findPath(root, q)

	var limit int
	if len(pPath) < len(qPath) {
		limit = len(pPath)
	} else {
		limit = len(qPath)
	}
	var previous = pPath[0]
	var checked = []TreeNode{}
	for i := 0; i < limit; i++ {
		if pPath[i].Val != qPath[i].Val {
			return &previous
		}
		checked = append(checked, pPath[i])
		previous = pPath[i]
	}

	return &checked[len(checked)-1]
}
