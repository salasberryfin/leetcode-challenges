package main

import "testing"

func TestCase1(t *testing.T) {
	s := "1001"
	expected := 2
	if result := minChanges(s); result != expected {
		t.Errorf("TestCase1 failed: expected %d, got %d", expected, result)
	}
}

func TestCase2(t *testing.T) {
	s := "10"
	expected := 1
	if result := minChanges(s); result != expected {
		t.Errorf("TestCase1 failed: expected %d, got %d", expected, result)
	}
}

func TestCase3(t *testing.T) {
	s := "0000"
	expected := 0
	if result := minChanges(s); result != expected {
		t.Errorf("TestCase1 failed: expected %d, got %d", expected, result)
	}
}
