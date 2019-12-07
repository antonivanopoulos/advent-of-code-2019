package main

import "testing"

func TestOrbit(t *testing.T) {
	for _, test := range orbitTests {
		if actual := Orbit(test.input); actual != test.expected {
			t.Errorf("Orbit(%v) expected %f, Actual %f", test.input, test.expected, actual)
		}
	}
}

func TestOrbitTwo(t *testing.T) {
	for _, test := range orbitTwoTests {
		if actual, _ := OrbitTwo(test.input); actual != test.expected {
			t.Errorf("OrbitTwo(%v) expected %f, Actual %f", test.input, test.expected, actual)
		}
	}
}
