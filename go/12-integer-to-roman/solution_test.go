package main

import "testing"

func TestExample1(t *testing.T) {
	num := 3
	expected := "III"
	result := intToRoman(num)

	if result != expected {
		t.Fatalf("Expected %s but got %s\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	num := 58
	expected := "LVIII"
	result := intToRoman(num)

	if result != expected {
		t.Fatalf("Expected %s but got %s\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	num := 1994
	expected := "MCMXCIV"
	result := intToRoman(num)

	if result != expected {
		t.Fatalf("Expected %s but got %s\n", expected, result)
	}
}

func TestExample4(t *testing.T) {
	num := 20
	expected := "XX"
	result := intToRoman(num)

	if result != expected {
		t.Fatalf("Expected %s but got %s\n", expected, result)
	}
}
