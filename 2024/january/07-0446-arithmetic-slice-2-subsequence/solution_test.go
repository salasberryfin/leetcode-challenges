package main

import "testing"

func TestExample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	output := 3
	result := numberOfArithmeticSlices(nums)
	if result != output {
		t.Errorf("Expected %d, but it was %d instead.", output, result)
	}
}
