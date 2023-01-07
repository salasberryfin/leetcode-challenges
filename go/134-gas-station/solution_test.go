package main

import "testing"

func TestExample1(t *testing.T) {
	gas := []int{
		1, 2, 3, 4, 5,
	}
	cost := []int{
		3, 4, 5, 1, 2,
	}

	expected := 3
	result := canCompleteCircuit(gas, cost)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	gas := []int{
		2, 3, 4,
	}
	cost := []int{
		3, 4, 3,
	}

	expected := -1
	result := canCompleteCircuit(gas, cost)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	gas := []int{
		5, 8, 2, 8,
	}
	cost := []int{
		6, 5, 6, 6,
	}

	expected := 3
	result := canCompleteCircuit(gas, cost)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample4(t *testing.T) {
	gas := []int{
		3, 1, 1,
	}
	cost := []int{
		1, 2, 2,
	}

	expected := 0
	result := canCompleteCircuit(gas, cost)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}
