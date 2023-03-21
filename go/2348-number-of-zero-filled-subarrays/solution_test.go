package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{
		1, 3, 0, 0, 2, 0, 0, 4,
	}
	output := int64(6)
	result := zeroFilledSubarray(nums)
	if output != result {
		t.Fatalf("zeroFilledSubarray(%v) = %d, expected %d\n", nums, result, output)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{
		0, 0, 0, 2, 0, 0,
	}
	output := int64(9)
	result := zeroFilledSubarray(nums)
	if output != result {
		t.Fatalf("zeroFilledSubarray(%v) = %d, expected %d\n", nums, result, output)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{
		2, 10, 2019,
	}
	output := int64(0)
	result := zeroFilledSubarray(nums)
	if output != result {
		t.Fatalf("zeroFilledSubarray(%v) = %d, expected %d\n", nums, result, output)
	}
}
