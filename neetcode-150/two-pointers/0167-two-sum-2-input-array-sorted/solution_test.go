package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	numbers := []int{
		2, 7, 11, 15,
	}
	target := 9
	output := []int{1, 2}
	result := twoSum(numbers, target)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	numbers := []int{
		2, 3, 4,
	}
	target := 6
	output := []int{1, 3}
	result := twoSum(numbers, target)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	numbers := []int{
		-1, 0,
	}
	target := -1
	output := []int{1, 2}
	result := twoSum(numbers, target)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
