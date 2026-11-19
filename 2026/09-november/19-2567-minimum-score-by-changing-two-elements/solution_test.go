package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{1, 4, 7, 8, 5}
	output := 3

	result := minimizeSum(nums)
	if result != output {
		t.Errorf("TestCase2 failed: expected %d, got %d", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 4, 3}
	output := 0

	result := minimizeSum(nums)
	if result != output {
		t.Errorf("TestCase2 failed: expected %d, got %d", output, result)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{31, 25, 72, 79, 74, 65}
	output := 14

	result := minimizeSum(nums)
	if result != output {
		t.Errorf("TestCase2 failed: expected %d, got %d", output, result)
	}
}
