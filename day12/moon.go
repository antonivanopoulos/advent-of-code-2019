package main

import (
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
	z int
}

type Velocity struct {
	x int
	y int
	z int
}

type Moon struct {
	coordinate Coordinate
	velocity   Velocity
}

func ParseInput(input []string) []Moon {
	moons := make([]Moon, len(input))
	for i, line := range input {
		slices := strings.Split(line, ",")

		xSlice := strings.TrimSpace(slices[0])
		x, _ := strconv.Atoi(xSlice[3:])

		ySlice := strings.TrimSpace(slices[1])
		y, _ := strconv.Atoi(ySlice[2:])

		zSlice := strings.TrimSpace(slices[2])
		z, _ := strconv.Atoi(zSlice[2 : len(zSlice)-1])

		moons[i] = Moon{Coordinate{x, y, z}, Velocity{0, 0, 0}}
	}

	return moons
}

func Step(input []Moon) []Moon {
	velocityChanges := make([][]Velocity, len(input))

	for i := 0; i < len(input); i++ {
		a := input[i]

		for j := i + 1; j < len(input); j++ {
			b := input[j]

			dAX, dAY, dAZ := a.DiffVelocity(b)
			velocityChanges[i] = append(velocityChanges[i], Velocity{dAX, dAY, dAZ})

			dBX, dBY, dBZ := b.DiffVelocity(a)
			velocityChanges[j] = append(velocityChanges[j], Velocity{dBX, dBY, dBZ})
		}
	}

	// Apply gravity changes
	for i, changes := range velocityChanges {
		for _, change := range changes {
			moon := input[i]
			input[i] = moon.ApplyGravity(change)
		}

		// Apply velocity
		moon := input[i]
		input[i] = moon.ApplyVelocity()
	}

	return input
}

func (a Moon) DiffVelocity(b Moon) (int, int, int) {
	diff := func(a int, b int) int {
		if b > a {
			return 1
		} else if b < a {
			return -1
		} else {
			return 0
		}
	}

	return diff(a.coordinate.x, b.coordinate.x),
		diff(a.coordinate.y, b.coordinate.y),
		diff(a.coordinate.z, b.coordinate.z)
}

func (a Moon) ApplyGravity(change Velocity) Moon {
	velocity := Velocity{a.velocity.x + change.x, a.velocity.y + change.y, a.velocity.z + change.z}
	return Moon{a.coordinate, velocity}
}

func (a Moon) ApplyVelocity() Moon {
	coordinate := Coordinate{
		a.coordinate.x + a.velocity.x,
		a.coordinate.y + a.velocity.y,
		a.coordinate.z + a.velocity.z,
	}

	return Moon{coordinate, a.velocity}
}

func (a Moon) CalculateEnergy() float64 {
	potentialEnergy := math.Abs(float64(a.coordinate.x)) + math.Abs(float64(a.coordinate.y)) + math.Abs(float64(a.coordinate.z))
	kineticEnergy := math.Abs(float64(a.velocity.x)) + math.Abs(float64(a.velocity.y)) + math.Abs(float64(a.velocity.z))

	return potentialEnergy * kineticEnergy
}
