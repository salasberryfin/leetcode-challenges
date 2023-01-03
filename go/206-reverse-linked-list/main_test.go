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
	expected := []int{5, 4, 3, 2, 1}
	list := reverseList(&input)
	result := printLinkedList(list)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	input := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	expected := []int{2, 1}
	list := reverseList(&input)
	result := printLinkedList(list)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	input := ListNode{}
	expected := []int{}
	list := reverseList(&input)
	result := printLinkedList(list)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
