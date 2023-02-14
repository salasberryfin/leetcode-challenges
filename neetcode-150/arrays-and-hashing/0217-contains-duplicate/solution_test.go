package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{
		1, 2, 3, 1,
	}
	output := true
	result := containsDuplicate(nums)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		1, 2, 3, 4,
	}
	output := false
	result := containsDuplicate(nums)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{
		1, 1, 1, 3, 3, 4, 3, 2, 4, 2,
	}
	output := true
	result := containsDuplicate(nums)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
