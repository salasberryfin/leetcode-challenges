package main

import "testing"

func TestExample1(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}
	start := 3
	output := 4
	result := amountOfTime(root, start)
	if result != output {
		t.Errorf("amountOfTime(root, start) = %d; want %d", result, output)
	}
}
