package main

import "testing"

func TestCase1(t *testing.T) {
	piles := []int{
		3, 6, 7, 11,
	}
	h := 8
	output := 4
	result := minEatingSpeed(piles, h)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	piles := []int{
		30, 11, 23, 4, 20,
	}
	h := 5
	output := 30
	result := minEatingSpeed(piles, h)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	piles := []int{
		30, 11, 23, 4, 20,
	}
	h := 6
	output := 23
	result := minEatingSpeed(piles, h)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
