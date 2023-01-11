package main

import "testing"

func TestExample1(t *testing.T) {
	n := 7
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 6},
	}
	hasApple := []bool{
		false,
		false,
		true,
		false,
		true,
		true,
		false,
	}
	expected := 8
	result := minTime(n, edges, hasApple)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	n := 7
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 6},
	}
	hasApple := []bool{
		false,
		false,
		true,
		false,
		false,
		true,
		false,
	}
	expected := 6
	result := minTime(n, edges, hasApple)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	n := 7
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 6},
	}
	hasApple := []bool{
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}
	expected := 0
	result := minTime(n, edges, hasApple)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}
