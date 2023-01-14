package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	s1, s2 := "parker", "morris"
	baseStr := "parser"
	expected := "makkek"
	result := smallestEquivalentString(s1, s2, baseStr)
	if result != expected {
		t.Fatalf("Expected %s but got %s\n", expected, result)
	}
}
