package main

import "testing"

func TestExample1(t *testing.T) {
	p := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	q := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
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
	result := isSameTree(&p, &q)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	p := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	q := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	expected := false
	result := isSameTree(&p, &q)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	p := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	q := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	expected := false
	result := isSameTree(&p, &q)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
