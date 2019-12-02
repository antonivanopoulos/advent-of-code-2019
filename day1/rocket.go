package main

// ModuleFuel calculates the fuel required given a rockets mass
func ModuleFuel(mass int) int {
	return (mass / 3) - 2
}
