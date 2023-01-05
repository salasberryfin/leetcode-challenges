package main

import "testing"

func TestExample1(t *testing.T) {
	prices := []int{
		7, 1, 5, 3, 6, 4,
	}
	expected := 5
	result := maxProfit(prices)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	prices := []int{
		7, 6, 4, 3, 1,
	}
	expected := 0
	result := maxProfit(prices)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	prices := []int{
		2, 4, 1,
	}
	expected := 2
	result := maxProfit(prices)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample4(t *testing.T) {
	prices := []int{
		2, 11, 1, 4, 7,
	}
	expected := 9
	result := maxProfit(prices)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample5(t *testing.T) {
	prices := []int{
		2, 1, 2, 0, 1,
	}
	expected := 1
	result := maxProfit(prices)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}
