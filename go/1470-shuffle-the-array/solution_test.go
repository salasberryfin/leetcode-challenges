package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{
		2, 5, 1, 3, 4, 7,
	}
	n := 3
	output := []int{
		2, 3, 5, 4, 1, 7,
	}
	result := shuffle(nums, n)
	if !reflect.DeepEqual(result, output) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		1, 2, 3, 4, 4, 3, 2, 1,
	}
	n := 4
	output := []int{
		1, 4, 2, 3, 3, 2, 4, 1,
	}
	result := shuffle(nums, n)
	if !reflect.DeepEqual(result, output) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{
		1, 1, 2, 2,
	}
	n := 2
	output := []int{
		1, 2, 1, 2,
	}
	result := shuffle(nums, n)
	if !reflect.DeepEqual(result, output) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
