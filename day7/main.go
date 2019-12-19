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

	program := make([]int, len(opcodes))
	copy(program, opcodes)

	output := Amplify(0, 0, program, make([]bool, 5))
	fmt.Printf("Max Output: %d\n", output)
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
