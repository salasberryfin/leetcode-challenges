package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	s := "bbbab"
	output := 4
	result := longestPalindromeSubseq(s)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	s := "cbbd"
	output := 2
	result := longestPalindromeSubseq(s)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
