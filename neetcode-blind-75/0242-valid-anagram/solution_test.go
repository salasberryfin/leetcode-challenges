package main

import "testing"

func TestCase1(test *testing.T) {
	s, t := "anagram", "nagaram"
	output := true
	result := isAnagram(s, t)
	if output != result {
		test.Fatalf("Expected %v but got %v\n", output, result)
	}
}
