package main

// ModuleFuel calculates the fuel required given a rockets mass
func ModuleFuel(mass int) int {
	fuel := (mass / 3) - 2

	if fuel > 0 {
		return fuel + ModuleFuel(fuel)
	}

	return 0
}
