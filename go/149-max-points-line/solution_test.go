package main

import "testing"

//func TestInLine1(t *testing.T) {
//	start, end := []int{1, 4}, []int{3, 2}
//	point := []int{2, 3}
//	expected := true
//	result := isInLine(point, start, end)
//
//	if expected != result {
//		t.Fatalf("Expected %v but got %v\n", expected, result)
//	}
//}
//
//func TestInLine2(t *testing.T) {
//	start, end := []int{1, 4}, []int{3, 2}
//	point := []int{1, 1}
//	expected := false
//	result := isInLine(point, start, end)
//
//	if expected != result {
//		t.Fatalf("Expected %v but got %v\n", expected, result)
//	}
//}
//
//func TestSlope1(t *testing.T) {
//	start, end := []int{1, 4}, []int{3, 2}
//	expected := -1
//	result := getSlope(start, end)
//
//	if expected != result {
//		t.Fatalf("Expected %v but got %v\n", expected, result)
//	}
//}
//
//func TestSlope2(t *testing.T) {
//	start, end := []int{1, 1}, []int{1, 4}
//	expected := 1000000
//	result := getSlope(start, end)
//
//	if expected != result {
//		t.Fatalf("Expected %v but got %v\n", expected, result)
//	}
//}
//
//func TestSlope3(t *testing.T) {
//	start, end := []int{1, 1}, []int{1, -4}
//	expected := -1000000
//	result := getSlope(start, end)
//
//	if expected != result {
//		t.Fatalf("Expected %v but got %v\n", expected, result)
//	}
//}

func TestExample1(t *testing.T) {
	points := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	expected := 3
	result := maxPoints(points)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	points := [][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}
	expected := 4
	result := maxPoints(points)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	points := [][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
		{1, 2},
	}
	expected := 4
	result := maxPoints(points)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample4(t *testing.T) {
	points := [][]int{
		{1, 0},
		{0, 0},
	}
	expected := 2
	result := maxPoints(points)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample5(t *testing.T) {
	points := [][]int{
		{0, 0},
		{4, 5},
		{7, 8},
		{8, 9},
		{5, 6},
		{3, 4},
		{1, 1},
	}
	expected := 5
	result := maxPoints(points)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample6(t *testing.T) {
	points := [][]int{
		{0, 0},
	}
	expected := 1
	result := maxPoints(points)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample7(t *testing.T) {
	points := [][]int{
		{0, 1},
		{0, 0},
		{0, 4},
		{0, -2},
		{0, -1},
		{0, 3},
		{0, -4},
	}
	expected := 7
	result := maxPoints(points)

	if expected != result {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
