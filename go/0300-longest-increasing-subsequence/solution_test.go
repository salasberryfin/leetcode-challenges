package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{
		10, 9, 2, 5, 3, 7, 101, 18,
	}
	output := 4
	result := lengthOfLIS(nums)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		0, 1, 0, 3, 2, 3,
	}
	output := 4
	result := lengthOfLIS(nums)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{
		7, 7, 7, 7, 7, 7, 7,
	}
	output := 1
	result := lengthOfLIS(nums)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
