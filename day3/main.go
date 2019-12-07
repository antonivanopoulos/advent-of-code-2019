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
	wires, err := ReadInputFile("./input.txt")
	check(err)

	var distance = ManhattenWiringDistance(wires)
	var delay = ManhattenWiringDelay(wires)
	fmt.Printf("Min Distance: %f\n", distance)
	fmt.Printf("Min Delay: %d\n", delay)
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
