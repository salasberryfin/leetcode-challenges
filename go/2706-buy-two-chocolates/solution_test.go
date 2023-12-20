package main

import (
	"testing"
)

func TestExample1(t *testing.T) {
	prices := []int{1, 2, 2}
	money := 3
	expected := 0
	result := buyChoco(prices, money)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestExample2(t *testing.T) {
	prices := []int{3, 2, 3}
	money := 3
	expected := 3
	result := buyChoco(prices, money)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestExample3(t *testing.T) {
	prices := []int{98, 54, 6, 34, 66, 63, 52, 39}
	money := 62
	expected := 22
	result := buyChoco(prices, money)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
