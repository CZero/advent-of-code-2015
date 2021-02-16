package main

import "testing"

var tests = []struct {
	n        string // input
	expected string // expected result
}{
	{"ugknbfddgicrmopn", "nice"},
	{"aaa", "nice"},
	{"jchzalrnumimnmhp", "naughty"},
	{"haegwjzuvuyypxyu", "naughty"},
	{"dvszwmarrgswjxmb", "naughty"},
}

func TestNaughtyOrNice(t *testing.T) {
	for _, test := range tests {
		actual := NaughtyOrNice(test.n)
		if actual != test.expected {
			t.Errorf("Input %-50s: expected %s, actual %s", test.n, test.expected, actual)
		}
	}
}
