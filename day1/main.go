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
	masses, err := ReadFrequencyFile("./input.txt")
	check(err)

	var total = 0
	for _, mass := range masses {
		total += ModuleFuel(mass)
	}

	fmt.Printf("Total Fuel: %d\n", total)
}

// ReadFrequencyFile does stuff
func ReadFrequencyFile(path string) (masses []int, err error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(dat), "\n")
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
