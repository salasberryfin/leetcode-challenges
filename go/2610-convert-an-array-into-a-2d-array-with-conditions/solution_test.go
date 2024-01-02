package main

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	nums := []int{1, 3, 4, 1, 2, 3, 1}
	expected := [][]int{{1, 3, 4, 2}, {1, 3}, {1}}

	result := findMatrix(nums)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := [][]int{{4, 3, 2, 1}}

	result := findMatrix(nums)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
