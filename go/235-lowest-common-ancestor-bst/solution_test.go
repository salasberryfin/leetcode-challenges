package main

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	root := TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   0,
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
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	p, q := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}, TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	expected := 6
	result := lowestCommonAncestor(&root, &p, &q)
	fmt.Printf("Expected %v and got %v\n", expected, result)

	//if expected != result {
	//	t.Fatalf("Expected %v but got %v\n", expected, result)
	//}
}

func TestExample2(t *testing.T) {
	root := TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   0,
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
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	p, q := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}, TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	expected := 2
	result := lowestCommonAncestor(&root, &p, &q)
	fmt.Printf("Expected %v and got %v\n", expected, result)

	//if expected != result {
	//	t.Fatalf("Expected %v but got %v\n", expected, result)
	//}
}

func TestExample3(t *testing.T) {
	root := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	p, q := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}, TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	expected := 2
	result := lowestCommonAncestor(&root, &p, &q)
	fmt.Printf("Expected %v and got %v\n", expected, result)

	//if expected != result {
	//	t.Fatalf("Expected %v but got %v\n", expected, result)
	//}
}

func TestExample4(t *testing.T) {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	p, q := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}, TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	expected := 3
	result := lowestCommonAncestor(&root, &p, &q)
	fmt.Printf("Expected %v and got %v\n", expected, result)

	//if expected != result {
	//	t.Fatalf("Expected %v but got %v\n", expected, result)
	//}
}
