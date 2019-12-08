package main

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestIntcode(t *testing.T) {
	for _, test := range intcodeTests {

		output := captureOutput(func() {
			if actual := Intcode(test.input, test.program); !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Intcode(%d, %d) expected %d, Actual %d", test.input, test.program, test.expected, actual)
			}
		})

		if output != test.output {
			t.Errorf("Intcode(%d, %d) expected output %s, Actual %s", test.input, test.program, test.output, output)
		}
	}
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
