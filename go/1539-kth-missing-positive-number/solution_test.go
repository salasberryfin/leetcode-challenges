package main

import "testing"

func TestCase1(t *testing.T) {
	arr := []int{
		2, 3, 4, 7, 11,
	}
	k := 5
	output := 9
	result := findKthPositive(arr, k)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
func TestCase2(t *testing.T) {
	arr := []int{
		1, 2, 3, 4,
	}
	k := 2
	output := 6
	result := findKthPositive(arr, k)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
