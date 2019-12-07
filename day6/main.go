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
	paths, err := ReadInputFile("./input.txt")
	check(err)

	var totalOrbits = Orbit(paths)
	fmt.Printf("Min Distance: %f\n", totalOrbits)

	var totalTransfers, _ = OrbitTwo(paths)
	fmt.Printf("Total Transfers between YOU and SAN: %f\n", totalTransfers)
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
