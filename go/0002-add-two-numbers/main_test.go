package main

import (
	"reflect"
	"testing"
)

func treeValuesToSlice(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}

	return result
}

func TestExample1(t *testing.T) {
	inputa := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	inputb := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	expected := []int{7, 0, 8}
	result := addTwoNumbers(&inputa, &inputb)

	resultSlice := treeValuesToSlice(result)

	if !reflect.DeepEqual(expected, resultSlice) {
		t.Fatalf("Expected %v but got %v\n", expected, resultSlice)
	}
}

func TestExample2(t *testing.T) {
	inputa := ListNode{
		Val:  0,
		Next: nil,
	}
	inputb := ListNode{
		Val:  0,
		Next: nil,
	}

	expected := []int{0}
	result := addTwoNumbers(&inputa, &inputb)

	resultSlice := treeValuesToSlice(result)

	if !reflect.DeepEqual(expected, resultSlice) {
		t.Fatalf("Expected %v but got %v\n", expected, resultSlice)
	}
}

func TestExample3(t *testing.T) {
	inputa := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	inputb := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}
	result := addTwoNumbers(&inputa, &inputb)

	resultSlice := treeValuesToSlice(result)

	if !reflect.DeepEqual(expected, resultSlice) {
		t.Fatalf("Expected %v but got %v\n", expected, resultSlice)
	}
}

func TestExample4(t *testing.T) {
	inputa := ListNode{
		Val:  0,
		Next: nil,
	}
	inputb := ListNode{
		Val:  1,
		Next: nil,
	}

	expected := []int{1}
	result := addTwoNumbers(&inputa, &inputb)

	resultSlice := treeValuesToSlice(result)

	if !reflect.DeepEqual(expected, resultSlice) {
		t.Fatalf("Expected %v but got %v\n", expected, resultSlice)
	}
}

func TestExample5(t *testing.T) {
	inputa := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}
	inputb := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	expected := []int{7, 0, 4, 0, 1}
	result := addTwoNumbers(&inputa, &inputb)

	resultSlice := treeValuesToSlice(result)

	if !reflect.DeepEqual(expected, resultSlice) {
		t.Fatalf("Expected %v but got %v\n", expected, resultSlice)
	}
}
