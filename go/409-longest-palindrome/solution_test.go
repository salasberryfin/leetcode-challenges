package main

import "testing"

func TestExample1(t *testing.T) {
	s := "abccccdd"
	expected := 7
	result := longestPalindrome(s)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	s := "a"
	expected := 1
	result := longestPalindrome(s)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	s := "aAAbbbccdaaaaaa"
	//"aaaAbcdcbAaaa"
	expected := 13
	result := longestPalindrome(s)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample4(t *testing.T) {
	s := ""
	for i := 0; i < 2000; i++ {
		s += "a"
	}
	expected := 2000
	result := longestPalindrome(s)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}
