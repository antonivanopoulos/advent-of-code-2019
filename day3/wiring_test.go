package main

import "testing"

func TestManhattenWiringDistance(t *testing.T) {
	for _, test := range manhattentWiringDistanceTests {
		if actual := ManhattenWiringDistance(test.input); actual != test.expected {
			t.Errorf("ManhattenWiringDistance(%v) expected %f, Actual %f", test.input, test.expected, actual)
		}
	}
}

func TestManhattenWiringDelay(t *testing.T) {
	for _, test := range manhattentWiringDelayTests {
		if actual := ManhattenWiringDelay(test.input); actual != test.expected {
			t.Errorf("ManhattenWiringDelay(%v) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}
