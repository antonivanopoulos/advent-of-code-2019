package main

import (
	"strconv"
	"strings"
)

func CleanSignal(signal string, patterns [][]int, n int, max int) string {
	if n >= max {
		return signal
	}

	output := make([]string, len(signal))
	for i := 0; i < len(signal); i++ {
		pattern := patterns[i]

		total := 0
		for j, char := range []rune(signal) {
			d := (int(char) - '0')
			total += d * pattern[j]
		}
		result := Abs(total % 10)
		output[i] = strconv.Itoa(result)
	}

	return CleanSignal(strings.Join(output, ""), patterns, n+1, max)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func BuildPattern(length int, n int) []int {
	var pattern []int

	base := []int{0, 1, 0, -1}

	for len(pattern) <= length+1 {
		for i := 0; i < n; i++ {
			pattern = append(pattern, base[0])

			if len(pattern) >= length+1 {
				break
			}
		}

		base = append(base[1:], base[0])
	}

	return pattern[1 : len(pattern)-1]
}
