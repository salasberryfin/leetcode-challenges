package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{
		1, 1, 2, 3, 3, 4, 4, 8, 8,
	}
	output := 2
	result := singleNonDuplicate(nums)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		3, 3, 7, 7, 10, 11, 11,
	}
	output := 10
	result := singleNonDuplicate(nums)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{
		1, 1, 2,
	}
	output := 2
	result := singleNonDuplicate(nums)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
