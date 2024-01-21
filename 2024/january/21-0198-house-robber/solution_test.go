package main

import "testing"

func TestExample1(t *testing.T) {
	input := []int{1, 2, 3, 1}
	expected := 4

	result := rob(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
