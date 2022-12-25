package main

import (
	"strings"
	"testing"
)

func TestLeetcodeInputExample1(t *testing.T) {
	input := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Right: nil,
			Left:  nil,
		},
		Left: &TreeNode{
			Val:   2,
			Right: nil,
			Left: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
		},
	}

	expected := "1(2(4))(3)"

	result := tree2str(&input)

	if !strings.EqualFold(result, expected) {
		t.Fatalf("tree2str() = %q - expected %#q\n", result, expected)
	}

}

func TestLeetcodeInputExample2(t *testing.T) {
	input := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Right: nil,
			Left:  nil,
		},
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	}

	expected := "1(2()(4))(3)"

	result := tree2str(&input)

	if !strings.EqualFold(result, expected) {
		t.Fatalf("tree2str() = %q - expected %#q\n", result, expected)
	}

}
