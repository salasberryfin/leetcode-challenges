package main

import (
	"log"
	"reflect"
	"testing"
)

//func TestCase1(t *testing.T) {
//	s := "aab"
//	output := [][]string{
//		{"a", "a", "b"},
//		{"aa", "b"},
//	}
//	result := partition(s)
//	if !reflect.DeepEqual(output, result) {
//		log.Fatalf("Expected %s but got %s\n", output, result)
//	}
//}
//
//func TestCase2(t *testing.T) {
//	s := "a"
//	output := [][]string{
//		{"a"},
//	}
//	result := partition(s)
//	if !reflect.DeepEqual(output, result) {
//		log.Fatalf("Expected %s but got %s\n", output, result)
//	}
//}

func TestCase3(t *testing.T) {
	s := "cbbbcc"
	output := [][]string{
		{"c", "b", "b", "b", "c", "c"},
		{"c", "b", "b", "b", "cc"},
		{"c", "b", "bb", "c", "c"},
		{"c", "b", "bb", "cc"},
		{"c", "bb", "b", "c", "c"},
		{"c", "bb", "b", "cc"},
		{"c", "bbb", "c", "c"},
		{"c", "bbb", "cc"},
		{"cbbbc", "c"},
	}
	//result := [][]string{
	//	{"c", "b", "b", "b", "c", "c"},
	//	{"c", "b", "b", "b", "cc"},
	//	{"c", "b", "bb", "c", "c"},
	//	{"c", "b", "bb", "cc"},
	//	{"c", "bb", "b", "c", "c"},
	//	{"c", "bb", "b", "cc"},
	//	{"c", "bbb", "c", "c"},
	//	{"c", "bbb", "cc"},
	//	{"cbbbc", "c"},
	//}
	result := partition(s)
	if !reflect.DeepEqual(output, result) {
		log.Fatalf("Expected %s but got %s\n", output, result)
	}
}

//func TestPalindrome1(t *testing.T) {
//	s := "arbra"
//	output := true
//	result := isPalindrome(s)
//	if output != result {
//		log.Fatalf("isPalindrome(s string) failed: expected %v but got %v\n", output, result)
//	}
//}
//
//func TestPalindrome2(t *testing.T) {
//	s := "bbb"
//	output := true
//	result := isPalindrome(s)
//	if output != result {
//		log.Fatalf("isPalindrome(s string) failed: expected %v but got %v\n", output, result)
//	}
//}
//
//func TestPalindrome3(t *testing.T) {
//	s := "arrara"
//	output := false
//	result := isPalindrome(s)
//	if output != result {
//		log.Fatalf("isPalindrome(s string) failed: expected %v but got %v\n", output, result)
//	}
//}
