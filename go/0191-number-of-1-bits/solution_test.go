package main

import "testing"

func TestCase1(t *testing.T) {
	n := uint32(00000000000000000000000000001011)
	output := 3
	result := hammingWeight(n)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
