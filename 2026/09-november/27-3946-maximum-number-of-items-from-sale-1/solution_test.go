package main

import "testing"

func TestCase1(t *testing.T) {
	items := [][]int{{6, 2}, {2, 6}, {3, 4}}
	budget := 9
	expected := 4
	result := maximumSaleItems(items, budget)
	if expected != result {
		t.Errorf("TestCase1 failed: expected %d, got %d", expected, result)
	}
}

func TestCase2(t *testing.T) {
	items := [][]int{{2, 4}, {3, 2}, {4, 1}, {6, 4}, {12, 4}}
	budget := 8
	expected := 10
	result := maximumSaleItems(items, budget)
	if expected != result {
		t.Errorf("TestCase2 failed: expected %d, got %d", expected, result)
	}
}
