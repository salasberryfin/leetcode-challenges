package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {

	root := TreeNode{
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

	expected := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}

	result := levelOrder(&root)
	fmt.Printf("Result %v\n", result)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
