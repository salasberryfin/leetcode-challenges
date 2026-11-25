package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{3, 3, 3}
	expected := 3
	result := smallestUniqueSubarray(nums)
	if result != expected {
		t.Errorf("TestCase1 failed: expected %d, got %d", expected, result)
	}
}
