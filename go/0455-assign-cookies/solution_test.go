package main

import "testing"

func TestExample1(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	output := 1
	if findContentChildren(g, s) != output {
		t.Fail()
	}
}

func TestExample2(t *testing.T) {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	output := 2
	if findContentChildren(g, s) != output {
		t.Fail()
	}
}

func TestExample3(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{3}
	output := 1
	if findContentChildren(g, s) != output {
		t.Fail()
	}
}

func TestExample4(t *testing.T) {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	output := 2
	if findContentChildren(g, s) != output {
		t.Fail()
	}
}
