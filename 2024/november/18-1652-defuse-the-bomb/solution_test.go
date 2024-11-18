package main

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	nums := []int{5, 7, 1, 4}
	k := 3
	output := []int{12, 10, 16, 13}
	result := decrypt(nums, k)
	if !reflect.DeepEqual(output, result) {
		t.Errorf("Expected %d, but it was %d instead.", output, result)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{5, 7, 1, 4}
	k := 0
	output := []int{0, 0, 0, 0}
	result := decrypt(nums, k)
	if !reflect.DeepEqual(output, result) {
		t.Errorf("Expected %d, but it was %d instead.", output, result)
	}
}

func TestExample3(t *testing.T) {
	nums := []int{2, 4, 9, 3}
	k := -2
	output := []int{12, 5, 6, 13}
	result := decrypt(nums, k)
	if !reflect.DeepEqual(output, result) {
		t.Errorf("Expected %d, but it was %d instead.", output, result)
	}
}
