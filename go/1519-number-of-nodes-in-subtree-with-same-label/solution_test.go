package main

import (
	"reflect"
	"testing"
)

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
	labels := "abaedcd"
	expected := []int{
		2, 1, 1, 1, 1, 1, 1,
	}
	result := countSubTrees(n, edges, labels)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample2(t *testing.T) {
	n := 7
	edges := [][]int{
		{0, 1},
		{1, 2},
		{0, 3},
	}
	labels := "bbbb"
	expected := []int{
		4, 2, 1, 1,
	}
	result := countSubTrees(n, edges, labels)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestExample3(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 2},
		{0, 3},
		{1, 2},
	}
	labels := "aeed"
	expected := []int{
		1, 1, 2, 1,
	}
	result := countSubTrees(n, edges, labels)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
