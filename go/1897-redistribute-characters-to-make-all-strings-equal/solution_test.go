package main

import "testing"

func TestExample1(t *testing.T) {
	words := []string{"abc", "aabc", "bc"}
	output := true
	result := makeEqual(words)
	if result != output {
		t.Errorf("Expected %v, but got %v", output, result)
	}
}

func TestExample2(t *testing.T) {
	words := []string{"ab", "a"}
	output := false
	result := makeEqual(words)
	if result != output {
		t.Errorf("Expected %v, but got %v", output, result)
	}
}
