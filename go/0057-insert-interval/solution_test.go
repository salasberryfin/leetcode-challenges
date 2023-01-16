package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	intervals := [][]int{
		{1, 3},
		{6, 9},
	}
	newInterval := []int{
		2, 5,
	}
	expected := [][]int{
		{1, 5},
		{6, 9},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase2(t *testing.T) {
	intervals := [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	newInterval := []int{
		4, 8,
	}
	expected := [][]int{
		{1, 2},
		{3, 10},
		{12, 16},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase3(t *testing.T) {
	intervals := [][]int{}
	newInterval := []int{
		5, 7,
	}
	expected := [][]int{
		{5, 7},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase4(t *testing.T) {
	intervals := [][]int{
		{1, 5},
	}
	newInterval := []int{
		2, 7,
	}
	expected := [][]int{
		{1, 7},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase5(t *testing.T) {
	intervals := [][]int{
		{1, 5},
	}
	newInterval := []int{
		6, 8,
	}
	expected := [][]int{
		{1, 5},
		{6, 8},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase6(t *testing.T) {
	intervals := [][]int{
		{1, 5},
	}
	newInterval := []int{
		2, 3,
	}
	expected := [][]int{
		{1, 5},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase7(t *testing.T) {
	intervals := [][]int{
		{1, 5},
	}
	newInterval := []int{
		0, 3,
	}
	expected := [][]int{
		{0, 5},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase8(t *testing.T) {
	intervals := [][]int{
		{1, 5},
	}
	newInterval := []int{
		0, 0,
	}
	expected := [][]int{
		{0, 0},
		{1, 5},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase9(t *testing.T) {
	intervals := [][]int{
		{2, 5},
		{6, 7},
		{8, 9},
	}
	newInterval := []int{
		0, 10,
	}
	expected := [][]int{
		{0, 10},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase10(t *testing.T) {
	intervals := [][]int{
		{1, 5},
		{6, 8},
	}
	newInterval := []int{
		5, 6,
	}
	expected := [][]int{
		{1, 8},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}

func TestCase11(t *testing.T) {
	intervals := [][]int{
		{2, 6},
		{7, 9},
	}
	newInterval := []int{
		15, 18,
	}
	expected := [][]int{
		{2, 6},
		{7, 9},
		{15, 18},
	}
	result := insert(intervals, newInterval)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
