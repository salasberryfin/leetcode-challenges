package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{
		1, 8, 8, 9,
	}
	target := 9
	output := [][]int{
		{1, 8},
		{1, 8},
	}

	result := twoSumWithRepetition(input, target)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
