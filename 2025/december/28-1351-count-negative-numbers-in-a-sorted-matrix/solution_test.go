package main

import "testing"

func TestExample1(t *testing.T) {
	grid := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	expected := 8
	if result := countNegatives(grid); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestExample2(t *testing.T) {
	grid := [][]int{
		{3, 2},
		{1, 0},
	}
	expected := 0
	if result := countNegatives(grid); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
