package main

import "testing"

func TestExample1(t *testing.T) {
	root := TreeNode{
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
	expected := true
	result := isValidBST(&root)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	root := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	expected := false
	result := isValidBST(&root)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	root := TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   -1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	expected := true
	result := isValidBST(&root)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample4(t *testing.T) {
	root := TreeNode{
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
	expected := false
	result := isValidBST(&root)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
