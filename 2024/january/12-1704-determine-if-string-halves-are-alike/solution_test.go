package main

import "testing"

func TestExample1(t *testing.T) {
	s := "book"
	output := true
	result := halvesAreAlike(s)
	if result != output {
		t.Errorf("Expected %v, got %v", output, result)
	}
}

func TestExample2(t *testing.T) {
	s := "textbook"
	output := false
	result := halvesAreAlike(s)
	if result != output {
		t.Errorf("Expected %v, got %v", output, result)
	}
}
