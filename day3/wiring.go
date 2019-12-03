package main

import (
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x float64
	y float64
}

// ManhattenWiring stuff
func ManhattenWiring(input []string) float64 {
	var paths [][]Coordinate

	for _, wire := range input {
		loc := Coordinate{0, 0}
		segments := strings.Split(wire, ",")
		var path []Coordinate

		for _, segment := range segments {
			direction := segment[0]
			distance, _ := strconv.Atoi(string(segment[1:]))

			for index := 0; index < distance; index++ {
				switch string(direction) {
				case "U":
					loc = Coordinate{loc.x, loc.y + 1}
				case "R":
					loc = Coordinate{loc.x + 1, loc.y}
				case "D":
					loc = Coordinate{loc.x, loc.y - 1}
				case "L":
					loc = Coordinate{loc.x - 1, loc.y}
				}

				path = append(path, loc)
			}
		}

		paths = append(paths, path)
	}

	pathA := paths[0]
	pathB := paths[1]
	var minDistance float64 = 9999999

	m := make(map[Coordinate]bool)

	for _, item := range pathA {
		m[item] = true
	}

	for _, item := range pathB {
		if _, ok := m[item]; ok {
			distance := math.Abs(item.x) + math.Abs(item.y)
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance
}
