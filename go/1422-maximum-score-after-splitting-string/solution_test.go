package main

import "testing"

func TestExample1(t *testing.T) {
	input := "011101"
	output := 5
	result := maxScore(input)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}

func TestExample2(t *testing.T) {
	input := "00111"
	output := 5
	result := maxScore(input)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}

func TestExample3(t *testing.T) {
	input := "1111"
	output := 3
	result := maxScore(input)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}

func TestExample4(t *testing.T) {
	input := "00"
	output := 1
	result := maxScore(input)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}
