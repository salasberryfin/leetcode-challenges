package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numberOfOnes(n int) int {
	var result int
	for n > 0 {
		result += n & 1
		n >>= 1
	}
	return result
}

func preorder(node *TreeNode, cnt int) int {
	if node == nil {
		return 0
	}
	cnt ^= (1 << node.Val)
	if node.Left == nil && node.Right == nil && numberOfOnes(cnt) <= 1 {
		return 1
	}
	return preorder(node.Left, cnt) + preorder(node.Right, cnt)
}

func pseudoPalindromicPaths(root *TreeNode) int {
	return preorder(root, 0)
}
