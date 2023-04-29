package main

import "testing"

func TestCase1(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	output := true
	result := isPalindrome(s)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	s := "race a car"
	output := false
	result := isPalindrome(s)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	s := "a9oo9a"
	output := true
	result := isPalindrome(s)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
