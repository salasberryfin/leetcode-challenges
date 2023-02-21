package main

import "testing"

func TestCase1(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	output := 1
	result := numIslands(grid)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
