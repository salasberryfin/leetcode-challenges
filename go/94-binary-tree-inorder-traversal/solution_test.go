package main

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	root := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Printf("%v", root)
}
