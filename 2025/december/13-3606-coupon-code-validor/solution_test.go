package main

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	code := []string{"SAVE20", "", "PHARMA5", "SAVE@20"}
	businessLine := []string{"restaurant", "grocery", "pharmacy", "restaurant"}
	isActive := []bool{true, true, true, true}
	output := []string{"PHARMA5", "SAVE20"}
	result := validateCoupons(code, businessLine, isActive)
	if !reflect.DeepEqual(output, result) {
		t.Errorf("Expected %v, but it was %v instead.", output, result)
	}
}

func TestExample2(t *testing.T) {
	code := []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"}
	businessLine := []string{"grocery", "electronics", "invalid"}
	isActive := []bool{false, true, true}
	output := []string{"ELECTRONICS_50"}
	result := validateCoupons(code, businessLine, isActive)
	if !reflect.DeepEqual(output, result) {
		t.Errorf("Expected %v, but it was %v instead.", output, result)
	}
}
