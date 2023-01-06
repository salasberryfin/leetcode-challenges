package main

import "testing"

func TestExample1(t *testing.T) {
	costs := []int{
		1, 3, 2, 4, 1,
	}
	coins := 7
	expected := 4
	result := maxIceCream(costs, coins)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	costs := []int{
		10, 6, 8, 7, 7, 8,
	}
	coins := 5
	expected := 0
	result := maxIceCream(costs, coins)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	costs := []int{
		1, 6, 3, 1, 2, 5,
	}
	coins := 20
	expected := 6
	result := maxIceCream(costs, coins)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
