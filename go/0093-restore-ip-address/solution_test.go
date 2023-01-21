package main

import (
	"log"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "25525511135"
	output := []string{
		"255.255.11.135",
		"255.255.111.35",
	}
	result := restoreIpAddress(s)
	if !reflect.DeepEqual(output, result) {
		log.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	s := "0000"
	output := []string{
		"0.0.0.0",
	}
	result := restoreIpAddress(s)
	if !reflect.DeepEqual(output, result) {
		log.Fatalf("Expected %v but got %v\n", output, result)
	}
}

func TestCase3(t *testing.T) {
	s := "101023"
	output := []string{
		"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3",
	}
	result := restoreIpAddress(s)
	if !reflect.DeepEqual(output, result) {
		log.Fatalf("Expected %v but got %v\n", output, result)
	}
}
