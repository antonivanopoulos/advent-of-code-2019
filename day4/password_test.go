package main

import "testing"

func TestPossiblePassword(t *testing.T) {
	for _, test := range possiblePasswordTests {
		if actual := PossiblePassword(test.input); actual != test.expected {
			t.Errorf("PossiblePassword(%s) expected %t, Actual %t", test.input, test.expected, actual)
		}
	}
}
