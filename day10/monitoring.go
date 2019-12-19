package main

import (
	"fmt"
	"math"
	"sort"
)

type Coordinate struct {
	x int
	y int
}

type Location struct {
	coord Coordinate
	count int
}

type Target struct {
	coord    Coordinate
	angle    float64
	distance float64
}

func FindIdealStationLocation(input []string) Location {
	los := make(map[Coordinate][]float64)

	for y1 := 0; y1 < len(input); y1++ {
		for x1 := 0; x1 < len(input[y1]); x1++ {
			if input[y1][x1] == '#' {
				for y2 := y1; y2 < len(input); y2++ {
					sx2 := 0
					if y2 == y1 {
						sx2 = x1 + 1
					}

					for x2 := sx2; x2 < len(input[y2]); x2++ {
						if rune(input[y2][x2]) == '#' {

							seen := false
							deg := CalculateAngle(x1, y1, x2, y2)
							for _, angle := range los[Coordinate{x1, y1}] {
								if angle == deg {
									seen = true
									break
								}
							}

							if !seen {
								los[Coordinate{x1, y1}] = append(los[Coordinate{x1, y1}], deg)
								los[Coordinate{x2, y2}] = append(los[Coordinate{x2, y2}], CalculateAngle(x2, y2, x1, y1))
							}
						}
					}
				}
			}
		}
	}

	mostSeen := 0
	var mostSeenCoord Coordinate
	for coord, seen := range los {
		if len(seen) > mostSeen {
			mostSeen = len(seen)
			mostSeenCoord = coord
		}
	}

	return Location{mostSeenCoord, mostSeen}
}

func BlastLazer(input [][]rune, location Location, score int) int {
	// deg := 90

	los := make(map[float64]Target)

	for y1 := 0; y1 < len(input); y1++ {
		for x1 := 0; x1 < len(input[y1]); x1++ {
			if input[y1][x1] == '#' {
				if x1 == location.coord.x && y1 == location.coord.y {
					continue
				}

				angle := CalculateAngle(x1, y1, location.coord.x, location.coord.y)
				distance := CalculateDistance(x1, y1, location.coord.x, location.coord.y)

				if target, ok := los[angle]; ok {
					if angle == target.angle && distance < target.distance {
						los[angle] = Target{Coordinate{x1, y1}, angle, distance}
					}
				} else {
					los[angle] = Target{Coordinate{x1, y1}, angle, distance}
				}
			}
		}
	}

	var targets []Target
	for _, target := range los {
		targets = append(targets, target)
	}

	sort.Slice(targets, func(i, j int) bool {
		return targets[i].angle > targets[j].angle
	})

	for _, target := range targets {
		score++

		if score == 199 {
			fmt.Println(target)
			return score
		}

		input[target.coord.y][target.coord.x] = '.'
	}

	return BlastLazer(input, location, score)
}

func CalculateAngle(x1 int, y1 int, x2 int, y2 int) float64 {
	deltaX := float64(x2 - x1)
	deltaY := float64(y2 - y1)

	angle := (math.Pi / 2) - math.Atan2(deltaY, deltaX)
	if angle < 0 {
		angle += math.Pi * 2
	}

	return angle
}

func CalculateDistance(x1 int, y1 int, x2 int, y2 int) float64 {
	deltaX := float64(x2 - x1)
	deltaY := float64(y2 - y1)

	return math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
}
