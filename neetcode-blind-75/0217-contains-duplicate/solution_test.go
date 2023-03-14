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
