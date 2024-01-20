package main

import "testing"

func TestExample1(t *testing.T) {
	arr := []int{3, 1, 2, 4}
	output := 17
	result := sumSubarrayMins(arr)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}

func TestExample2(t *testing.T) {
	arr := []int{11, 81, 94, 43, 3}
	output := 444
	result := sumSubarrayMins(arr)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}
