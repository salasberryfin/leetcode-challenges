package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{
		-1, 0, 3, 5, 9, 12,
	}
	target := 9
	output := 4
	result := search(nums, target)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		-1, 0, 3, 5, 9, 12,
	}
	target := 2
	output := -1
	result := search(nums, target)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
