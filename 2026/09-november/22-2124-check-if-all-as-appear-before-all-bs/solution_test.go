package main

import "testing"

func TestCase1(t *testing.T) {
	input := "aaabbb"
	output := true
	result := checkString(input)
	if result != output {
		t.Errorf("TestCase1 failed: expected %v, got %v", output, result)
	}
}

func TestCase2(t *testing.T) {
	input := "abab"
	output := false
	result := checkString(input)
	if result != output {
		t.Errorf("TestCase1 failed: expected %v, got %v", output, result)
	}
}

func TestCase3(t *testing.T) {
	input := "bbb"
	output := true
	result := checkString(input)
	if result != output {
		t.Errorf("TestCase1 failed: expected %v, got %v", output, result)
	}
}
