package main

import "testing"

func TestExample1(t *testing.T) {
	m := 2
	n := 2
	maxMove := 2
	startRow := 0
	startColumn := 0
	want := 6
	got := findPaths(m, n, maxMove, startRow, startColumn)
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestExample2(t *testing.T) {
	m := 1
	n := 3
	maxMove := 3
	startRow := 0
	startColumn := 1
	want := 12
	got := findPaths(m, n, maxMove, startRow, startColumn)
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
