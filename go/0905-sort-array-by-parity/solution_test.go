package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{
		3, 1, 2, 4,
	}
	output := []int{
		2, 4, 3, 1,
	}
	result := sortArrayByParity(nums)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		0,
	}
	output := []int{
		0,
	}
	result := sortArrayByParity(nums)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
