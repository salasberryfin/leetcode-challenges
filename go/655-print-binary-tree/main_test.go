package main

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	input := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	result := printTree(&input)

	expected := [][]string{
		{"", "1", ""},
		{"2", "", ""},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	input := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
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

	result := printTree(&input)

	expected := [][]string{
		{"", "", "", "1", "", "", ""},
		{"", "2", "", "", "", "3", ""},
		{"", "", "4", "", "", "", ""},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
