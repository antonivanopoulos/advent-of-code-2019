package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x    float64
	y    float64
	step int
}

// ManhattenWiring stuff
func ManhattenWiring(input []string) [][]Coordinate {
	var paths [][]Coordinate

	for _, wire := range input {
		loc := Coordinate{0, 0, 0}
		segments := strings.Split(wire, ",")
		var path []Coordinate

		for _, segment := range segments {
			direction := segment[0]
			distance, _ := strconv.Atoi(string(segment[1:]))

			for index := 0; index < distance; index++ {
				switch string(direction) {
				case "U":
					loc = Coordinate{loc.x, loc.y + 1, loc.step + 1}
				case "R":
					loc = Coordinate{loc.x + 1, loc.y, loc.step + 1}
				case "D":
					loc = Coordinate{loc.x, loc.y - 1, loc.step + 1}
				case "L":
					loc = Coordinate{loc.x - 1, loc.y, loc.step + 1}
				}

				path = append(path, loc)
			}
		}

		paths = append(paths, path)
	}

	return paths
}

func ManhattenWiringDistance(input []string) float64 {
	paths := ManhattenWiring(input)
	pathA := paths[0]
	pathB := paths[1]
	var minDistance float64 = 9999999

	m := make(map[string]bool)

	for _, item := range pathA {
		m[item.key()] = true
	}

	for _, item := range pathB {
		if _, ok := m[item.key()]; ok {
			distance := math.Abs(item.x) + math.Abs(item.y)
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance
}

func ManhattenWiringDelay(input []string) int {
	paths := ManhattenWiring(input)
	pathA := paths[0]
	pathB := paths[1]
	var minDelay int = 9999999

	m := make(map[string]int)

	for _, item := range pathA {
		m[item.key()] = item.step
	}

	for _, item := range pathB {
		if _, ok := m[item.key()]; ok {
			delay := m[item.key()] + item.step
			if delay < minDelay {
				minDelay = delay
			}
		}
	}

	return minDelay
}

func (coord Coordinate) key() string {
	return fmt.Sprintf("%fx%f", coord.x, coord.y)
}
