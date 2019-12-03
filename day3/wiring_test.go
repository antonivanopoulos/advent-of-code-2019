package main

import "testing"

func TestManhattenWiring(t *testing.T) {
	for _, test := range manhattentWiringTests {
		if actual := ManhattenWiring(test.input); actual != test.expected {
			t.Errorf("ManhattenWiring(%v) expected %f, Actual %f", test.input, test.expected, actual)
		}
	}
}
