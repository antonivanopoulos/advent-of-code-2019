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

	moons := ParseInput(input)
	origMoons := make([]Moon, len(moons))
	for i, m := range moons {
		var cpy Moon
		cpy.coordinate = m.coordinate
		cpy.velocity = m.velocity

		origMoons[i] = cpy
	}

	moonsCopy := make([]Moon, len(moons))
	copy(moonsCopy, moons)

	// Part A
	for index := 0; index < 10; index++ {
		moonsCopy = Step(moonsCopy)
	}

	totalEnergy := float64(0)
	for _, moon := range moonsCopy {
		totalEnergy += moon.CalculateEnergy()
	}
	fmt.Println(moonsCopy)
	fmt.Printf("Sum of total energy: %f\n", totalEnergy)

	// Part B
	var txRepeat int = 0
	var tyRepeat int = 0
	var tzRepeat int = 0
	for t := 1; ; t++ {
		moons = Step(moons)

		var xAxisNotReapting = false
		var yAxisNotReapting = false
		var zAxisNotReapting = false

		for i, moon := range moons {
			if moon.coordinate.x != origMoons[i].coordinate.x || moon.velocity.x != origMoons[i].velocity.x {
				xAxisNotReapting = true
			}

			if moon.coordinate.y != origMoons[i].coordinate.y || moon.velocity.y != origMoons[i].velocity.y {
				yAxisNotReapting = true
			}

			if moon.coordinate.z != origMoons[i].coordinate.z || moon.velocity.z != origMoons[i].velocity.z {
				zAxisNotReapting = true
			}
		}

		if !xAxisNotReapting && txRepeat == 0 {
			txRepeat = t
		}

		if !yAxisNotReapting && tyRepeat == 0 {
			tyRepeat = t
		}

		if !zAxisNotReapting && tzRepeat == 0 {
			tzRepeat = t
		}

		if txRepeat*tyRepeat*tzRepeat > 0 {
			fmt.Println(txRepeat, tyRepeat, tzRepeat)

			result := txRepeat * tyRepeat / gcd(txRepeat, tyRepeat)
			result = result * tzRepeat / gcd(result, tzRepeat)
			fmt.Println(result)
			return
		}
	}
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

func gcd(p, q int) int {
	if q == 0 {
		return p
	}

	r := p % q
	return gcd(q, r)
}
