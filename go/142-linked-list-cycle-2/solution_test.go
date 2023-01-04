package main

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	input := ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	expected := []int{2, 0, -4}
	result := detectCycle(&input)
	repr := printLinkedList(result)

	if !reflect.DeepEqual(expected, repr) {
		t.Fatalf("Expected %v but got %v\n", expected, repr)
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
	expected := []int{1, 2}
	result := detectCycle(&input)
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
	expected := []int{-1}
	result := detectCycle(&input)
	repr := printLinkedList(result)

	if !reflect.DeepEqual(expected, repr) {
		t.Fatalf("Expected %v but got %v\n", expected, repr)
	}
}
