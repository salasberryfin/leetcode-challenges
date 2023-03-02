package main

import "testing"

func TestCase1(t *testing.T) {
	chars := []byte{
		'a', 'a', 'b', 'b', 'c', 'c', 'c',
	}
	output := 6
	result := compress(chars)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	chars := []byte{
		'a',
	}
	output := 1
	result := compress(chars)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	chars := []byte{
		'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b',
	}
	output := 4
	result := compress(chars)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase4(t *testing.T) {
	chars := []byte{
		'a', 'a', 'a', 'b', 'b', 'a', 'a',
	}
	output := 6
	result := compress(chars)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase5(t *testing.T) {
	chars := []byte{
		'a', 'b', 'c',
	}
	output := 3
	result := compress(chars)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
