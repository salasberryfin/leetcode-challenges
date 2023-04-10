package main

import "testing"

func TestCase1(t *testing.T) {
	s := "()"
	output := true
	result := isValid(s)
	if output != result {
		t.Fatalf("isValid(%s) = %v, expected %v\n", s, result, output)
	}
}

func TestCase2(t *testing.T) {
	s := "()[]{}"
	output := true
	result := isValid(s)
	if output != result {
		t.Fatalf("isValid(%s) = %v, expected %v\n", s, result, output)
	}
}

func TestCase3(t *testing.T) {
	s := "(]"
	output := false
	result := isValid(s)
	if output != result {
		t.Fatalf("isValid(%s) = %v, expected %v\n", s, result, output)
	}
}
