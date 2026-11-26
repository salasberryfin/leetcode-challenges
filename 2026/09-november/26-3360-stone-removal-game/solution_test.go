package main

import "testing"

func TestCase1(t *testing.T) {
	n := 12
	expected := true
	result := canAliceWin(n)
	if result != expected {
		t.Errorf("TestCase1 failed: expected %v, got %v", expected, result)
	}
}

func TestCase2(t *testing.T) {
	n := 1
	expected := false
	result := canAliceWin(n)
	if result != expected {
		t.Errorf("TestCase2 failed: expected %v, got %v", expected, result)
	}
}
