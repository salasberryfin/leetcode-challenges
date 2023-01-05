package main

import "testing"

func TestExample1(t *testing.T) {
	input := 3
	expected := []string{""}
	result := generateParenthesis(input)

	t.Fatalf("Expected %v but got %v\n", expected, result)
}
