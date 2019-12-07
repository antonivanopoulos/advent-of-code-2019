package main

import (
	"strings"
)

type Path struct {
	orbitee string
	orbiter string
}

func Orbit(input []string) float64 {
	tree := make(map[string][]string)

	for _, str := range input {
		path := ConvertToPath(str)
		tree[path.orbitee] = append(tree[path.orbitee], path.orbiter)
	}

	total := CountPaths(tree, "COM", 0)
	return total
}

func OrbitTwo(input []string) (transfers float64, ok bool) {
	tree := make(map[string][]string)

	for _, str := range input {
		path := ConvertToPath(str)
		tree[path.orbitee] = append(tree[path.orbitee], path.orbiter)
	}

	var aParent, bParent string
	var aOk, bOk bool
	var aCount, bCount float64

	aParent, aOk = FindParent(tree, "YOU")
	aCount = 0
	for {
		bParent, bOk = FindParent(tree, "SAN")
		bCount = 0
		for {
			bParent, bOk = FindParent(tree, bParent)
			bCount++

			if !bOk {
				break
			}

			if aParent == bParent {
				var transfers float64 = aCount + bCount
				return transfers, true
			}
		}

		aParent, aOk = FindParent(tree, aParent)
		aCount++
		if !aOk {
			return 0, false
		}
	}
}

func ConvertToPath(input string) Path {
	split := strings.Split(input, ")")
	return Path{split[0], split[1]}
}

func CountPaths(tree map[string][]string, node string, depth float64) float64 {
	nodes := tree[node]
	total := depth

	for _, node := range nodes {
		total += CountPaths(tree, node, depth+1)
	}

	return total
}

func FindParent(tree map[string][]string, node string) (parent string, ok bool) {
	for k, v := range tree {
		for _, leaf := range v {
			if leaf == node {
				return k, true
			}
		}
	}

	return "", false
}
