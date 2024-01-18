package main

import (
	"testing"
)

func TestExample1(t *testing.T) {
	n := 2
	output := 2

	result := climbStairs(n)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}

func TestExample2(t *testing.T) {
	n := 3
	output := 3

	result := climbStairs(n)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}

func TestExample3(t *testing.T) {
	n := 5
	output := 8

	result := climbStairs(n)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}
