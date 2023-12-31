package main

import "testing"

func TestExample1(t *testing.T) {
	s := "aa"
	expected := 0
	actual := maxLengthBetweenEqualCharacters(s)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}

func TestExample2(t *testing.T) {
	s := "abca"
	expected := 2
	actual := maxLengthBetweenEqualCharacters(s)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}

func TestExample3(t *testing.T) {
	s := "cbzxy"
	expected := -1
	actual := maxLengthBetweenEqualCharacters(s)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}

func TestExample4(t *testing.T) {
	s := "cbzxyzuiopz"
	expected := 7
	actual := maxLengthBetweenEqualCharacters(s)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
