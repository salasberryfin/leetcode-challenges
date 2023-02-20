package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{
		1, 3, 5, 6,
	}
	target := 5
	output := 2
	result := searchInsert(nums, target)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		1, 3, 5, 6,
	}
	target := 2
	output := 1
	result := searchInsert(nums, target)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{
		1, 3, 5, 6,
	}
	target := 7
	output := 4
	result := searchInsert(nums, target)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
