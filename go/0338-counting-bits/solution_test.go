package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	n := 2
	output := []int{
		0, 1, 1,
	}
	result := countBits(n)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	n := 5
	output := []int{
		0, 1, 1, 2, 1, 2,
	}
	result := countBits(n)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	n := 0
	output := []int{
		0,
	}
	result := countBits(n)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("Expected %v but got %v\n", output, result)
	}
}
