package main

import "testing"

func TestCase1(t *testing.T) {
	intervals := []*Interval{
		{
			Start: 0,
			End:   30,
		},
		{
			Start: 5,
			End:   10,
		},
		{
			Start: 15,
			End:   20,
		},
	}
	output := 2
	result := minMeetingRooms(intervals)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}

func TestCase2(t *testing.T) {
	intervals := []*Interval{
		{
			Start: 2,
			End:   7,
		},
	}
	output := 1
	result := minMeetingRooms(intervals)
	if output != result {
		t.Fatalf("Expected %d but got %d\n", output, result)
	}
}
