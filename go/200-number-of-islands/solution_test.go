package main

import "testing"

func TestExample1(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	expected := 1
	result := numIslands(grid)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	grid := [][]byte{
		{'0', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	expected := 1
	result := numIslands(grid)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	expected := 2
	result := numIslands(grid)

	if expected != result {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

//func TestExample4(t *testing.T) {
//	grid := [][]byte{}
//	for i := 0; i < 200; i++ {
//		row := []byte{}
//		for j := 0; j < 200; j++ {
//			row = append(row, '1')
//		}
//		grid = append(grid, row)
//	}
//	expected := 2
//	result := numIslands(grid)
//
//	if expected != result {
//		t.Fatalf("Expected %d but got %d\n", expected, result)
//	}
//}
