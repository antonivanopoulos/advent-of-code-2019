package main

import (
	"fmt"
	"strconv"
)

func main() {
	total := 0
	for value := 183564; value <= 657474; value++ {
		str := strconv.Itoa(value)
		if PossiblePassword(str) {
			total++
		}
	}

	fmt.Printf("Total possible passwords: %d\n", total)
}
