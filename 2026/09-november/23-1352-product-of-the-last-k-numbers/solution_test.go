package main

import "testing"

func TestCase1(t *testing.T) {
	obj := Constructor()
	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)

	if got := obj.GetProduct(2); got != 20 {
		t.Errorf("GetProduct(2) = %d; want 20", got)
	}
}
