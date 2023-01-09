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
	result := preorderTraversal(&root)
	fmt.Printf("Result: %v\n", result)
}

func TestExample2(t *testing.T) {
	root := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
	}
	preorderTraversal(&root)
	result := preorderTraversal(&root)
	fmt.Printf("Result: %v\n", result)
}

func TestExample3(t *testing.T) {
	root := TreeNode{
		Val: 3,
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
	preorderTraversal(&root)
	result := preorderTraversal(&root)
	fmt.Printf("Result: %v\n", result)
}
