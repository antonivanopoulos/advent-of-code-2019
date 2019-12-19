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

	location := FindIdealStationLocation(input)
	fmt.Printf("Best location: %d, %d.. Count: %d\n", location.coord.x, location.coord.y, location.count)

	var lazerInput [][]rune
	for _, line := range input {
		lazerInput = append(lazerInput, []rune(line))
	}

	BlastLazer(lazerInput, location, 0)
}

// ReadInputFile reads in the given input file, and converts the supplied ops codes into an array of ints
func ReadInputFile(path string) (wires []string, err error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	return lines, nil
}
