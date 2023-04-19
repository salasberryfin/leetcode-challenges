package main

import "testing"

func TestCase1(t *testing.T) {
	root := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:  1,
						Left: nil,
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	output := 3
	result := longestZigZag(&root)
	if result != output {
		t.Fatalf("expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	output := 4
	result := longestZigZag(&root)
	if result != output {
		t.Fatalf("expected %d but got %d\n", output, result)
	}
}
