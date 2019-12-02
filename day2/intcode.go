package main

// Intcode runs through the list of provided integers and performs the operations
// represented by those codes (addition, multiplication. and halting)
func Intcode(input []int) []int {
	index := 0

	for index < len(input) {
		switch input[index] {
		case 1:
			input[input[index+3]] = input[input[index+1]] + input[input[index+2]]
			index += 4
		case 2:
			input[input[index+3]] = input[input[index+1]] * input[input[index+2]]
			index += 4
		case 99:
			return input
		}
	}

	return input
}
