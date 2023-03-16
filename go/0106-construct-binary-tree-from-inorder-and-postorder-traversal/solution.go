/*
*
Last value in postorder traversal is always the root node
First value in inorder traversal is always the leftmost node
*
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	inorderMap := map[int]int{}
	postorderMap := map[int]int{}
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
		postorderMap[postorder[i]] = i
	}

	last := len(postorder) - 1
	result := &TreeNode{Val: postorder[last]}
	mid := inorderMap[postorder[last]]
	result.Left = buildTree(inorder[:mid], postorder[:mid])
	result.Right = buildTree(inorder[mid+1:], postorder[mid:last])

	return result
}
