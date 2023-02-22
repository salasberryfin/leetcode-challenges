package main

import "testing"

func TestCase1(t *testing.T) {
	n := 2
	output := 2
	result := climbStairs(n)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	n := 3
	output := 3
	result := climbStairs(n)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
