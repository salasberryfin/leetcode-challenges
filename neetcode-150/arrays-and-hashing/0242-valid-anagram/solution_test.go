package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "anagram"
	ts := "nagaram"
	output := true
	result := isAnagram(s, ts)
	if output != result {
		log.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	s := "rat"
	ts := "car"
	output := false
	result := isAnagram(s, ts)
	if output != result {
		log.Fatalf("Expected %v but got %v\n", output, result)
	}
}
