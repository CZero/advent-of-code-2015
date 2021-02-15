package main

import "testing"

var tests = []struct {
	n        string // input
	expected int    // expected result
}{
	{">", 2},
	{"^>v<", 4},
	{"^v^v^v^v^v", 2},
}

func TestCountHousesOnDirections(t *testing.T) {
	for _, test := range tests {
		actual := CountHousesOnDirections(test.n)
		if actual != test.expected {
			t.Errorf("Input %-50q: expected %d, actual %d", test.n, test.expected, actual)
		}
	}
}
