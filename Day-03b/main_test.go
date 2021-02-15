package main

import "testing"

var tests = []struct {
	in       string // input
	expected int    // expected result
}{
	{"^>", 3},
	{"^>v<", 3},
	{"^v^v^v^v^v", 11},
}

func TestCountHousesOnDirections(t *testing.T) {
	for _, test := range tests {
		actual := CountHousesOnDirections(test.in)
		if actual != test.expected {
			t.Errorf("Input %-50q: expected %d, actual %d", test.in, test.expected, actual)
		}
	}
}
