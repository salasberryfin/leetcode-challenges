package main

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	input := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	expected := []int{3, 4, 5}
	result := middleNode(&input)
	repr := printLinkedList(result)

	if !reflect.DeepEqual(expected, repr) {
		t.Fatalf("Expected %v but got %v\n", expected, repr)
	}
}

func TestExample2(t *testing.T) {
	input := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	expected := []int{4, 5, 6}
	result := middleNode(&input)
	repr := printLinkedList(result)

	if !reflect.DeepEqual(expected, repr) {
		t.Fatalf("Expected %v but got %v\n", expected, repr)
	}
}

func TestExample3(t *testing.T) {
	input := ListNode{
		Val:  1,
		Next: nil,
	}
	expected := []int{1}
	result := middleNode(&input)
	repr := printLinkedList(result)

	if !reflect.DeepEqual(expected, repr) {
		t.Fatalf("Expected %v but got %v\n", expected, repr)
	}
}
