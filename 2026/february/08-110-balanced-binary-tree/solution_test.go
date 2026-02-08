package main

import "testing"

func TestExample1(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	output := true
	result := isBalanced(root)

	if result != output {
		t.Errorf("Expected %v, got %v", output, result)
	}
}

func TestExample2(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 4}

	output := false
	result := isBalanced(root)

	if result != output {
		t.Errorf("Expected %v, got %v", output, result)
	}
}
