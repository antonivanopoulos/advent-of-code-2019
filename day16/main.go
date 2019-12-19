package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, err := ReadInputFile("./input.txt")
	check(err)

	fmt.Println(input)

	n := 0
	patterns := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		pattern := BuildPattern(len(input), i+1)
		patterns[i] = pattern
	}

	signal := CleanSignal(input, patterns, n, 100)
	fmt.Println(signal[:8])
}

// ReadInputFile reads in the given input file, and converts the supplied ops codes into an array of ints
func ReadInputFile(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	input := strings.TrimSpace(string(dat))

	return input, nil
}
