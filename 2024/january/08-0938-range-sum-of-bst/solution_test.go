package main

import "testing"

func TestExample1(t *testing.T) {
	low, high := 7, 15
	// root = [10,5,15,3,7,null,18]
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  15,
			Left: nil,
			Right: &TreeNode{
				Val:   18,
				Left:  nil,
				Right: nil,
			},
		},
	}
	output := 32
	result := rangeSumBST(root, low, high)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}
