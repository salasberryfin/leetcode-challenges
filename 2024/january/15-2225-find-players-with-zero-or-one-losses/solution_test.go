package main

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	matches := [][]int{
		{1, 3},
		{2, 3},
		{3, 6},
		{5, 6},
		{5, 7},
		{4, 5},
		{4, 8},
		{4, 9},
		{10, 4},
		{10, 9},
	}

	expected := [][]int{
		{1, 2, 10},
		{4, 5, 7, 8},
	}

	result := findWinners(matches)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
