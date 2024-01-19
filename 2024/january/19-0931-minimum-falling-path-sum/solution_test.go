package main

import "testing"

func TestExample1(t *testing.T) {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	output := 13

	result := minFallingPathSum(matrix)
	if result != output {
		t.Errorf("Expected %d, got %d", output, result)
	}
}
