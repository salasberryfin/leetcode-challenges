package main

import (
	"log"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{
		4, 6, 7, 7,
	}
	expected := [][]int{
		{4, 6},
		{4, 6, 7},
		{4, 6, 7, 7},
		{4, 7},
		{4, 7, 7},
		{6, 7},
		{6, 7, 7},
		{7, 7},
	}
	result := findSubsequences(nums)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalf("Expected %v but got %v\n", expected, result)
	}
}
