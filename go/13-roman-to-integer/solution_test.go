package main

import "testing"

func TestExample1(t *testing.T) {
	s := "III"
	expected := 3
	result := romanToInteger(s)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	s := "LVIII"
	expected := 58
	result := romanToInteger(s)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	s := "MCMXCIV"
	expected := 1994
	result := romanToInteger(s)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample4(t *testing.T) {
	s := "XX"
	expected := 20
	result := romanToInteger(s)

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}
