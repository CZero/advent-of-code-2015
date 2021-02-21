package main

import "testing"

var tests = []struct {
	n        string // input
	expected string // expected result
}{
	{"qjhvhtzxzqqjkmpb", "nice"},
	{"xxyxx", "nice"},
	{"uurcxstgmygtbstg", "naughty"},
	{"ieodomkazucvgmuy", "naughty"},
}

func TestNaughtyOrNice(t *testing.T) {
	for _, test := range tests {
		actual := NaughtyOrNice(test.n)
		if actual != test.expected {
			t.Errorf("Input %-50s: expected %s, actual %s", test.n, test.expected, actual)
		}
	}
}

var testOne = []struct {
	n        string // input
	expected bool   // expected result
}{
	{"xyxy", true},
	{"aabcdefgaa", true},
	{"aaa", false},
	{"ewglkjwrgsadmnawe", false},
	{"erwxrsglkjwrgsadmwewe", true},
}

func TestTestOne(t *testing.T) {
	for _, test := range testOne {
		actual := TestOne(test.n)
		if actual != test.expected {
			t.Errorf("Input %-50s: expected %t, actual %t", test.n, test.expected, actual)
		}
	}
}

var testTwo = []struct {
	n        string // input
	expected bool   // expected result
}{
	{"xyx", true},
	{"abcdefeghi", true},
	{"aaa", true},
	{"dwegjsexe", true},
	{"aab", false},
	{"slgdslkghw", false},
}

func TestTestTwo(t *testing.T) {
	for _, test := range testTwo {
		actual := TestTwo(test.n)
		if actual != test.expected {
			t.Errorf("Input %-50s: expected %t, actual %t", test.n, test.expected, actual)
		}
	}
}
