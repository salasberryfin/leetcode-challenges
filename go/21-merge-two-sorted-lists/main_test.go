package main

import "testing"

func TestExample1(t *testing.T) {
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	result := mergeTwoLists(&list1, &list2)
	expected := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}

	if *result != expected {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
