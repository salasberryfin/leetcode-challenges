package main

import "testing"

func TestExample1(t *testing.T) {
	points := [][]int{
		{10, 16},
		{2, 8},
		{1, 6},
		{7, 12},
	}
	expected := 2
	result := findMinArrowShots(points)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	points := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	expected := 4
	result := findMinArrowShots(points)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	points := [][]int{}
	for i := 0; i < 10000; i++ {
		points = append(points, []int{i, i + 1})
	}
	expected := 2
	result := findMinArrowShots(points)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample4(t *testing.T) {
	points := [][]int{
		{1, 2},
	}
	expected := 1
	result := findMinArrowShots(points)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample5(t *testing.T) {
	points := [][]int{
		{3, 9},
		{7, 12},
		{3, 8},
		{6, 8},
		{9, 10},
		{2, 9},
		{0, 9},
		{3, 9},
		{0, 6},
		{2, 8},
	}
	expected := 2
	result := findMinArrowShots(points)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}
