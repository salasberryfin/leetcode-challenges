package main

import (
	"log"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	strs := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	output := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	result := groupAnagrams(strs)
	if !reflect.DeepEqual(output, result) {
		log.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	strs := []string{
		"",
	}
	output := [][]string{
		{""},
	}
	result := groupAnagrams(strs)
	if !reflect.DeepEqual(output, result) {
		log.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	strs := []string{
		"a",
	}
	output := [][]string{
		{"a"},
	}
	result := groupAnagrams(strs)
	if !reflect.DeepEqual(output, result) {
		log.Fatalf("Expected %v but got %v\n", output, result)
	}
}
