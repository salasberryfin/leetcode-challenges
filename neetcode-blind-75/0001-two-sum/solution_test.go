package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{
		2, 7, 11, 15,
	}
	target := 9
	output := []int{
		0, 1,
	}
	result := twoSum(nums, target)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
