package main

import "testing"

func TestExample1(t *testing.T) {
	s := "bab"
	tt := "aba"
	expected := 1

	result := minSteps(s, tt)
	if result != expected {
		t.Errorf("minSteps(%s, %s) returned %d but expected %d", s, tt, result, expected)
	}
}

func TestExample2(t *testing.T) {
	s := "leetcode"
	tt := "practice"
	expected := 5

	result := minSteps(s, tt)
	if result != expected {
		t.Errorf("minSteps(%s, %s) returned %d but expected %d", s, tt, result, expected)
	}
}

func TestExample3(t *testing.T) {
	s := "anagram"
	tt := "mangaar"
	expected := 0

	result := minSteps(s, tt)
	if result != expected {
		t.Errorf("minSteps(%s, %s) returned %d but expected %d", s, tt, result, expected)
	}
}
