package main

import "testing"

func TestCase1(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	output := 3
	result := diameterOfBinaryTree(root)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
