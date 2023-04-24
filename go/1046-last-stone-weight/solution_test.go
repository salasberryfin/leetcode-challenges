package main

import "testing"

func TestCase1(t *testing.T) {
	stones := []int{
		2, 7, 4, 1, 8, 1,
	}
	output := 1
	result := lastStoneWeight(stones)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	stones := []int{1}
	output := 1
	result := lastStoneWeight(stones)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
