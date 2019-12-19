package main

import (
	"testing"
)

func TestAmplify(t *testing.T) {
	for _, test := range amplificationTests {
		if actual := Amplify(test.input, 0, test.program, make([]bool, 5)); actual != test.expected {
			t.Errorf("Amplify(%d, %d, %d) expected %d, Actual %d", test.input, 0, test.program, test.expected, actual)
		}
	}
}
