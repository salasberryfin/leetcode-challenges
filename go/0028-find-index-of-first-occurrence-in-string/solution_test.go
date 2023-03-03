package main

import "testing"

func TestCase1(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	output := 0
	result := strStr(haystack, needle)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	haystack := "leetcode"
	needle := "leeto"
	output := -1
	result := strStr(haystack, needle)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
