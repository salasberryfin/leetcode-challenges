package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	n := 2
	trust := [][]int{
		{1, 2},
	}
	expected := 2
	result := findJudge(n, trust)
	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestCase2(t *testing.T) {
	n := 3
	trust := [][]int{
		{1, 3},
		{2, 3},
	}
	expected := 3
	result := findJudge(n, trust)
	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestCase3(t *testing.T) {
	n := 3
	trust := [][]int{
		{1, 3},
		{2, 3},
		{3, 1},
	}
	expected := -1
	result := findJudge(n, trust)
	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}
