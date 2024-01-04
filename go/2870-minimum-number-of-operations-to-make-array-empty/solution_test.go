package main

import "testing"

func TestExample1(t *testing.T) {
	nums := []int{2, 3, 3, 2, 2, 4, 2, 3, 4}
	expected := 4
	actual := minOperations(nums)
	if actual != expected {
		t.Errorf("TestExample1: expected %d, got %d", expected, actual)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{2, 1, 2, 2, 3, 3}
	expected := -1
	actual := minOperations(nums)
	if actual != expected {
		t.Errorf("TestExample1: expected %d, got %d", expected, actual)
	}
}

func TestExample3(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2, 3, 3}
	expected := 3
	actual := minOperations(nums)
	if actual != expected {
		t.Errorf("TestExample1: expected %d, got %d", expected, actual)
	}
}

func TestExample4(t *testing.T) {
	nums := []int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}
	expected := 7
	actual := minOperations(nums)
	if actual != expected {
		t.Errorf("TestExample1: expected %d, got %d", expected, actual)
	}
}
func TestExample5(t *testing.T) {
	nums := []int{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19}
	expected := 5
	actual := minOperations(nums)
	if actual != expected {
		t.Errorf("TestExample1: expected %d, got %d", expected, actual)
	}
}
