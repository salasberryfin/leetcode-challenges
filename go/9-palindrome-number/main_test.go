package main

import "testing"

func TestExample1(t *testing.T) {
	input := 121
	expected := true
	result := isPalindrome(input)

	if result != expected {
		t.Fatalf("Expected %v, got %v\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	input := -121
	expected := false
	result := isPalindrome(input)

	if result != expected {
		t.Fatalf("Expected %v, got %v\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	input := 10
	expected := false
	result := isPalindrome(input)

	if result != expected {
		t.Fatalf("Expected %v, got %v\n", expected, result)
	}
}
