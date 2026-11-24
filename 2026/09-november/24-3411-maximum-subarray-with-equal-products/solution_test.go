package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 1, 2, 1, 1, 1}
	expected := 5
	result := maxLength(nums)
	if result != expected {
		t.Errorf("TestCase1 failed: expected %d, got %d", expected, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{2, 3, 4, 5, 6}
	expected := 3
	result := maxLength(nums)
	if result != expected {
		t.Errorf("TestCase1 failed: expected %d, got %d", expected, result)
	}
}
