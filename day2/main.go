package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	opcodes, err := ReadInputFile("./input.txt")
	check(err)

	// Perform the state reset
	for noun := range [100]int{} {
		for verb := range [100]int{} {
			program := make([]int, len(opcodes))
			copy(program, opcodes)

			program[1] = noun
			program[2] = verb

			var result = Intcode(program)
			if result[0] == 19690720 {
				fmt.Printf("Result: %d\nNoun: %d\nVerb: %d", result[0], noun, verb)
				break
			}
		}
	}
}

// ReadInputFile reads in the given input file, and converts the supplied ops codes into an array of ints
func ReadInputFile(path string) (masses []int, err error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), ",")
	masses = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		masses = append(masses, n)
	}

	return masses, nil
}
