/*
One function to traverse both trees, if elements are different, return false
Else, keep going

Check if any of the trees is nil:
	- if one is nil but not the other: false
	- if both nil: true
Check if both tree.Val are equal:
	- if true: keep going
	- if false: return false
Go left for both trees
Go right for both trees
Check results for both subtrees:
	- If left == right -> return either of them
	- If left != right: at least one of them is dalse -> return false
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	if left == right {
		return left
	} else {
		return false
	}
}
