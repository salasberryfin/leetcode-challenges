package main

import "testing"

func TestCase2(t *testing.T) {
	s := "()[]{}"
	output := true
	result := isValid(s)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	s := "(]"
	output := false
	result := isValid(s)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase4(t *testing.T) {
	s := "(])"
	output := false
	result := isValid(s)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
