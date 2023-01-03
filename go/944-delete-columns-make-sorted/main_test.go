package main

import "testing"

func TestSolutionExample1(t *testing.T) {
	input := []string{"cba", "daf", "ghi"}
	expected := 1
	result := minDeletionSize(input)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestSolutionExample2(t *testing.T) {
	input := []string{"a", "b"}
	expected := 0
	result := minDeletionSize(input)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestSolutionExample3(t *testing.T) {
	input := []string{"zyx", "wvu", "tsr"}
	expected := 3
	result := minDeletionSize(input)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}
