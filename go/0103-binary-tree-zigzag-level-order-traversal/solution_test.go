package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
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
	output := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}
	result := zigzagLevelOrder(root)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
