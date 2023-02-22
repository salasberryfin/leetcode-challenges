package main

import "testing"

func TestCase1(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	output := true
	result := isValidBST(root)

	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
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
	output := false
	result := isValidBST(root)

	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 6,
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
	}
	output := false
	result := isValidBST(root)

	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
