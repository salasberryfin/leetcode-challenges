package main

import "testing"

func TestExample1(t *testing.T) {
	arr := []int{1, 2, 2, 1, 1, 3}
	output := true

	result := uniqueOccurrences(arr)

	if result != output {
		t.Errorf("uniqueOccurrences(%v) = %v, want %v", arr, result, output)
	}
}

func TestExample2(t *testing.T) {
	arr := []int{1, 2}
	output := false

	result := uniqueOccurrences(arr)

	if result != output {
		t.Errorf("uniqueOccurrences(%v) = %v, want %v", arr, result, output)
	}
}

func TestExample3(t *testing.T) {
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	output := true

	result := uniqueOccurrences(arr)

	if result != output {
		t.Errorf("uniqueOccurrences(%v) = %v, want %v", arr, result, output)
	}
}
