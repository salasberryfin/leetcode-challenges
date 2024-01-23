package main

import "testing"

func TestExample1(t *testing.T) {
	arr := []string{"un", "iq", "ue"}
	expected := 4
	actual := maxLength(arr)
	if actual != expected {
		t.Errorf("TestExample1: expected %v, actual %v", expected, actual)
	}
}

func TestExample2(t *testing.T) {
	arr := []string{"cha", "r", "act", "ers"}
	expected := 6
	actual := maxLength(arr)
	if actual != expected {
		t.Errorf("TestExample2: expected %v, actual %v", expected, actual)
	}
}
