package main

import "testing"

func TestExample1(t *testing.T) {
	nums := []int{
		1, 2, 0,
	}
	expected := 3
	result := firstMissingPositive(nums)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{
		3, 4, -1, 1,
	}
	expected := 2
	result := firstMissingPositive(nums)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	nums := []int{
		7, 8, 9, 11, 12,
	}
	expected := 1
	result := firstMissingPositive(nums)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample4(t *testing.T) {
	nums := []int{
		-30, 7, 8, 9, 11, 12, 10, 1, 2, 5,
	}
	expected := 3
	result := firstMissingPositive(nums)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample5(t *testing.T) {
	nums := []int{
		-30, 7, 8, 9, 3, 4, 6, 11, 12, 10, 1, 2, 5, -13,
	}
	expected := 13
	result := firstMissingPositive(nums)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample6(t *testing.T) {
	nums := []int{
		2,
	}
	expected := 1
	result := firstMissingPositive(nums)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample7(t *testing.T) {
	nums := []int{
		-5,
	}
	expected := 1
	result := firstMissingPositive(nums)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
