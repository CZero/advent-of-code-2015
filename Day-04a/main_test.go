package main

import "testing"

var tests = []struct {
	in       string // input
	expected int    // expected result
}{
	{"abcdef", 609043},
	{"pqrstuv", 1048970},
}

func Test(t *testing.T) {
	for _, test := range tests {
		actual := FindFirstCoin(test.in)
		if actual != test.expected {
			t.Errorf("Input %-50q: expected %d, actual %d", test.in, test.expected, actual)
		}
	}
}
