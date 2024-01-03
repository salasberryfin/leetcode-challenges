package main

import "testing"

func TestExample1(t *testing.T) {
	bank := []string{"011001", "000000", "010100", "001000"}
	output := 8
	if numberOfBeams(bank) != output {
		t.Errorf("Expected %d, got %d", output, numberOfBeams(bank))
	}
}
