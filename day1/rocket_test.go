package main

import "testing"

func TestModuleFuel(t *testing.T) {
	for _, test := range moduleFuelTests {
		if actual := ModuleFuel(test.input); actual != test.expected {
			t.Errorf("ModuleFuel(%d) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}
