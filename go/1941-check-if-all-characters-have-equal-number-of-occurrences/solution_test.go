package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	s := "abacbc"
	output := true
	result := areOccurrencesEqual(s)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
