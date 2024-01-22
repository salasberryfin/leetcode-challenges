package main

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	nums := []int{1, 2, 2, 4}
	expected := []int{2, 3}
	actual := findErrorNums(nums)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{1, 1}
	expected := []int{1, 2}
	actual := findErrorNums(nums)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestExample3(t *testing.T) {
	nums := []int{3, 2, 2}
	expected := []int{2, 1}
	actual := findErrorNums(nums)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestExample4(t *testing.T) {
	nums := []int{3, 2, 3, 4, 6, 5}
	expected := []int{3, 1}
	actual := findErrorNums(nums)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
