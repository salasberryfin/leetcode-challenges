package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{
		3, 0, 1,
	}
	output := 2
	result := missingNumber(nums)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		0, 1,
	}
	output := 2
	result := missingNumber(nums)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{
		9, 6, 4, 2, 3, 5, 7, 0, 1,
	}
	output := 8
	result := missingNumber(nums)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
