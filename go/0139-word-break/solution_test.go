package main

import "testing"

func TestCase1(t *testing.T) {
	s := "leetcode"
	wordDict := []string{
		"leet",
		"code",
	}
	output := true
	result := wordBreak(s, wordDict)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	s := "applepenapple"
	wordDict := []string{
		"apple",
		"pen",
	}
	output := true
	result := wordBreak(s, wordDict)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	s := "catsandog"
	wordDict := []string{
		"cats",
		"dog",
		"sand",
		"and",
		"cat",
	}
	output := false
	result := wordBreak(s, wordDict)
	if output != result {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
