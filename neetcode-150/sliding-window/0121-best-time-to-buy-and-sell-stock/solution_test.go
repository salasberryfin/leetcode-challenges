package main

import "testing"

func TestCase1(t *testing.T) {
	prices := []int{
		7, 1, 5, 3, 6, 4,
	}
	output := 5
	result := maxProfit(prices)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	prices := []int{
		7, 6, 4, 3, 1,
	}
	output := 0
	result := maxProfit(prices)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
