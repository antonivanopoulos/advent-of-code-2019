package main

import (
	"reflect"
	"testing"
)

func TestIntcode(t *testing.T) {
	for _, test := range intcodeTests {
		if actual := Intcode(test.input); !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Intcode(%d) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}
