package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{
		3, 4, 5, 1, 2,
	}
	output := 1
	result := findMin(nums)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		4, 5, 6, 7, 0, 1, 2,
	}
	output := 0
	result := findMin(nums)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{
		11, 13, 15, 17,
	}
	output := 11
	result := findMin(nums)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase4(t *testing.T) {
	nums := []int{
		2, 1,
	}
	output := 1
	result := findMin(nums)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase5(t *testing.T) {
	nums := []int{
		3, 1, 2,
	}
	output := 1
	result := findMin(nums)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
