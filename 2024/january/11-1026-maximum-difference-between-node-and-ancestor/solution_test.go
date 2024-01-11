package main

import "testing"

func TestExample1(t *testing.T) {
	root := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 1, Left: nil, Right: nil},
			Right: &TreeNode{
				Val:   6,
				Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
				Right: &TreeNode{Val: 7, Left: nil, Right: nil},
			},
		},
		Right: &TreeNode{
			Val:  10,
			Left: nil,
			Right: &TreeNode{
				Val:   14,
				Left:  &TreeNode{Val: 13, Left: nil, Right: nil},
				Right: nil,
			},
		},
	}

	output := 7
	result := maxAncestorDiff(root)

	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}
